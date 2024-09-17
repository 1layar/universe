package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	ttemplate "text/template"

	"github.com/1layar/universe/internal/email_service/internal/app/appconfig"
	"github.com/1layar/universe/internal/email_service/internal/dto"
	"github.com/1layar/universe/internal/email_service/model"
	"github.com/hibiken/asynq"
)

type EmailManager struct {
	messageService  *emailMessageService
	templateService *emailTemplateService
	eventService    *emailEventService
	agent           *emailAgent
	client          *asynq.Client
}

// define const interface here
var _ IEmailManager = (*EmailManager)(nil)

func NewEmailManager(
	client *asynq.Client,
	messageService *emailMessageService,
	agent *emailAgent,
	templateService *emailTemplateService,
	eventService *emailEventService,
) *EmailManager {
	return &EmailManager{
		messageService:  messageService,
		agent:           agent,
		client:          client,
		templateService: templateService,
		eventService:    eventService,
	}
}

// Compose implements EmailManager.
func (m *EmailManager) Compose(ctx context.Context, compose dto.Compose) (*model.EmailEvent, error) {
	agent, err := m.agent.account.GetAccountByCode(ctx, compose.Agent)

	if err != nil {
		return nil, err
	}

	// get template by code
	emailTemplate, err := m.templateService.GetEmailTemplateByCode(ctx, compose.Template)

	if err != nil {
		return nil, err
	}

	// union payload from compose and template
	// create io writer for subject
	payload := emailTemplate.Placeholders

	for k, v := range compose.Payload {
		payload[k] = v
	}

	// create io writer for subject
	subject, err := m.parseTemplate(emailTemplate.Subject, payload)
	if err != nil {
		return nil, err
	}

	// create io writer for body
	bodyHtml, err := m.parseTemplate(emailTemplate.HtmlContent, payload)
	if err != nil {
		return nil, err
	}

	// create io writer for body
	bodyText, err := m.parseTemplate(emailTemplate.TextContent, payload)
	if err != nil {
		return nil, err
	}

	emailMsgData := &model.EmailMessage{
		ToEmail:   compose.Email,
		AccountID: agent.ID,
		TextBody:  bodyText,
		HtmlBody:  bodyHtml,
		Subject:   subject,
	}
	// create email message
	err = m.messageService.CreateEmailMessage(ctx, emailMsgData)

	if err != nil {
		return nil, err
	}

	eventData := &model.EmailEvent{
		TemplateID: emailTemplate.ID,
		MessageID:  emailMsgData.ID,
		Payload:    compose.Payload,
		EventType:  compose.Event,
	}

	err = m.eventService.CreateEmailEvent(ctx, eventData)

	if err != nil {
		return nil, err
	}

	eventData.Message = emailMsgData
	eventData.Template = emailTemplate

	return eventData, nil
}

func (*EmailManager) parseTemplate(templ string, payload map[string]any) (string, error) {
	template := ttemplate.New("")

	res, err := template.Parse(templ)

	if err != nil {
		return "", err
	}

	subject := &bytes.Buffer{}

	err = res.Execute(subject, payload)

	if err != nil {
		return "", err
	}

	return subject.String(), nil
}

func (m *EmailManager) Send(messageID int) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(appconfig.EmailDeliveryPayload{MessageID: messageID})
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(appconfig.TypeEmailDelivery, payload)

	info, err := m.client.Enqueue(task)
	if err != nil {
		return nil, fmt.Errorf("could not schedule task: %w", err)
	}

	return info, nil
}

/*
This will call agent and send email
*/
func (m *EmailManager) ProcessEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	payload := &appconfig.EmailDeliveryPayload{}
	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}

	message, err := m.messageService.GetEmailMessageByID(ctx, payload.MessageID)

	if err != nil {
		return err
	}

	return m.agent.SendEmail(ctx, *message)
}
