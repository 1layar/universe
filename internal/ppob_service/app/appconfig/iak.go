package appconfig

type IakProductOperator string

type IakProductType string

type IakProductStatus string

const (
	Pulsa IakProductType = "pulsa"
	Data  IakProductType = "data"
)

const (
	Active   IakProductStatus = "active"
	Inactive IakProductStatus = "non active"
	All      IakProductStatus = "all"
)

const (
	PulsaAxis      IakProductOperator = "axis"
	PulsaIndosat   IakProductOperator = "indosat"
	PulsaSmartfren IakProductOperator = "smart"
	PulsaTelkom    IakProductOperator = "telkomsel"
	PulsaTri       IakProductOperator = "three"
	PulsaXixiGame  IakProductOperator = "xixi_games"
	PulsaXl        IakProductOperator = "xl"
	PulsaByU       IakProductOperator = "by.U"
)

type IakStatusCode string

const (
	CODE_00  IakStatusCode = "00"  // SUCCESS	Success	Transaction success.
	CODE_06  IakStatusCode = "06"  // TRANSACTION NOT FOUND	Failed	There is no transaction with your inputted ref_id. Please check again your inputted ref_id to find your transaction.
	CODE_07  IakStatusCode = "07"  // FAILED	Failed	Your current transaction has failed. Please try again.
	CODE_10  IakStatusCode = "10"  // REACH TOP UP LIMIT USING SAME DESTINATION NUMBER IN 1 DAY	Failed	Your current destination number top up request is reaching the limit on that day. Please try again tomorrow.
	CODE_12  IakStatusCode = "12"  // BALANCE MAXIMUM LIMIT EXCEEDED	Failed	-
	CODE_13  IakStatusCode = "13"  // CUSTOMER NUMBER BLOCKED	Failed	Your customer number (customer_id) has been blocked. You can change your customer number (customer_id) or contact our Customer Service.
	CODE_14  IakStatusCode = "14"  // INCORRECT DESTINATION NUMBER	Failed	Your customer_id that you’ve inputted isn’t a valid phone number. Please check again your customer_id.
	CODE_16  IakStatusCode = "16"  // NUMBER NOT MATCH WITH OPERATOR	Failed	Your phone number (customer_id) that you’ve inputted doesn’t match with your desired operator (product_code). Please check again your phone number or change your operator.
	CODE_17  IakStatusCode = "17"  // INSUFFICIENT DEPOSIT	Failed	Your current deposit is lower than the product_price you want to buy. You can add more money into your deposit by doing top up on iak.id deposit menu, or if you are in development mode, you can add your development deposit by clicking the + (plus) sign on development deposit menu (https://developer.iak.id/sandbox-report).
	CODE_18  IakStatusCode = "18"  // NUMBER NOT AVAILABLE	Failed	You can see all available E-SIM number by using E-SIM List API.
	CODE_19  IakStatusCode = "19"  // NUMBER IS ALREADY IN USE	Failed	Please select other E-SIM number.
	CODE_20  IakStatusCode = "20"  // CODE NOT FOUND	Failed	Your inputted product_code isn’t in the database. Check again your product_code, you can check product_code list by using Pricelist API.
	CODE_21  IakStatusCode = "21"  // NUMBER EXPIRED	Failed	Your phone number (customer_id) is expired. You can try other phone number.
	CODE_39  IakStatusCode = "39"  // PROCESS	Pending	Your current transaction is being processed, please wait until your transaction is fully processed. You can check the status by using check-status API or by receiving a callback (if you use callback).
	CODE_102 IakStatusCode = "102" // INVALID IP ADDRESS	Failed	Your IP address isn’t allowed to make a transaction. You can add your IP address to your allowed IP address list in https://developer.iak.id/prod-setting.
	CODE_106 IakStatusCode = "106" // PRODUCT IS TEMPORARILY OUT OF SERVICE	Failed	The product_code that you pick is in non-active status. You can retry your transaction with another product_code that has active status.
	CODE_107 IakStatusCode = "107" // ERROR IN XML FORMAT	Failed	The body format of your request isn’t correct or there is an error in your body (required, ajax error, etc). Please use the correct JSON or XML format corresponding to your request to API. You can see the required body request for each request in the API Documentation.
	CODE_110 IakStatusCode = "110" // SYSTEM UNDER MAINTENANCE	Failed	The system is currently under maintenance, you can try again later.
	CODE_117 IakStatusCode = "117" // PAGE NOT FOUND	Failed	The API URL that you want to hit is not found. Try checking your request URL for any typos or try other API URLs.
	CODE_121 IakStatusCode = "121" // MONTHLY TOP UP LIMIT EXCEEDED	Failed	This response code applies to OVO products.
	CODE_131 IakStatusCode = "131" // TOP UP REGION BLOCKED FOR PLAYER	Failed	Your current destination number top up request is blocked in that region. Please try again with a different destination number.
	CODE_132 IakStatusCode = "132" // PRODUCT CODE NOT ELIGIBLE DUE TO SUBSCRIBER LOCATION	Failed	Your inputted product_code isn’t eligible due to subscriber location. Please try again with different product_code.
	CODE_141 IakStatusCode = "141" // INVALID USER ID / ZONE ID / SERVER ID / ROLENAME	Failed	Your inputted user ID / Zone ID / Server ID / Role name isn’t valid. Please try again with another user ID / Zone ID / Server ID / Role name. You can check on Inquiry Game Server.
	CODE_142 IakStatusCode = "142" // INVALID USER ID	Failed	Your current destination number (user id) top up request is invalid. Please try again with a different destination number or try checking for typos in your field.
	CODE_201 IakStatusCode = "201" // UNDEFINED RESPONSE CODE	Pending	The received response code is undefined yet. Please contact our Customer Service.
	CODE_202 IakStatusCode = "202" // MAXIMUM 1 NUMBER 1 TIME IN 1 DAY	Failed	You can only top up to a phone number once in a day (based on your developer setting). If you want to allow more than one top up to a phone number, you can go to https://developer.iak.id/ then choose API Setting menu, you can turn on “Allow multiple transactions for the same number” in development or production settings.
	CODE_203 IakStatusCode = "203" // NUMBER IS TOO LONG	Failed	Your inputted customer ID is too long. Please check again your customer ID.
	CODE_204 IakStatusCode = "204" // WRONG AUTHENTICATION	Failed	Your sign (signature) field doesn’t contain the right key for your current request. Please check again your sign value.
	CODE_205 IakStatusCode = "205" // WRONG COMMAND	Failed	The command that you’ve inputted is not a valid command, try check your commands field for typos or try another command.
	CODE_206 IakStatusCode = "206" // THIS DESTINATION NUMBER HAS BEEN BLOCKED	Failed	The customer_id that you inputted is blocked or not in whitelist. You can unblock it by remove customer number blacklist in API Security menu blacklist (https://developer.iak.id/end-user-blacklist) or add customer number whitelist in API Security menu whitelist (https://developer.iak.id/end-user-whitelist) on developer.iak.id.
	CODE_207 IakStatusCode = "207" // MAXIMUM 1 NUMBER WITH ANY CODE 1 TIME IN 1 DAY	Failed	You’ve already done a transaction today. Please do another transaction tomorrow, or disable the high restriction setting in https://developer.iak.id/prod-setting.
)

type IakOperatorPrefix []string

var (
	INDOSAT   IakOperatorPrefix = []string{"0814", "0815", "0816", "0855", "0856", "0857", "0858"}
	XL        IakOperatorPrefix = []string{"0817", "0818", "0819", "0859", "0878", "0877"}
	AXIS      IakOperatorPrefix = []string{"0838", "0837", "0831", "0832"}
	TELKOMSEL IakOperatorPrefix = []string{"0812", "0813", "0852", "0853", "0821", "0823", "0822", "0851"}
	SMARTFREN IakOperatorPrefix = []string{"0881", "0882", "0883", "0884", " 0885", "0886", "0887", "0888"}
	THREE     IakOperatorPrefix = []string{"0896", "0897", "0898", "0899", "0895"}
	ByU       IakOperatorPrefix = []string{"085154", "085155", "085156", "085157", "085158"}
)

type IakPlnInqueryStatus string

const (
	Success IakPlnInqueryStatus = "1"
	Failed  IakPlnInqueryStatus = "2"
)

type GameCode string

const (
	GAME_103 GameCode = "103" //Mobile Legend		{userid}|{zoneid}
	GAME_127 GameCode = "127" //Ragnarok		{userid}|{serverid}
	GAME_130 GameCode = "130" //Point Blank		{userid}
	GAME_135 GameCode = "135" //Free Fire		{userid}
	GAME_136 GameCode = "136" //Speed Drifters		{userid}
	GAME_139 GameCode = "139" //Arena of Valor		{userid}
	GAME_140 GameCode = "140" //Bleach Mobile 3D		{rolename}|{userid}|{serverid}
	GAME_141 GameCode = "141" //Era of Celestial		{rolename}|{userid}|{serverid}
	GAME_142 GameCode = "142" //Dragon Nest		{rolename}|{serverid}
	GAME_146 GameCode = "146" //Call of Duty		{userid}
	GAME_150 GameCode = "150" //Marvel Super War		{userid}
	GAME_152 GameCode = "152" //Light of Thel:Glory of Cepheus		{userid}
	GAME_153 GameCode = "153" //Lords Mobile		{userid}
	GAME_154 GameCode = "154" //Life After		{userid}|{serverid}
	GAME_172 GameCode = "172" //Genshin Impact		{userid}|{serverid}
	GAME_176 GameCode = "176" //LoL Wild Rift		{userid}|{tag}
	GAME_230 GameCode = "230" //Heroes Evolved		{userid}|{serverid}
)

var CODE_TEMPLATE = map[GameCode]string{
	GAME_103: "{userid}|{zoneid}",
	GAME_127: "{userid}|{serverid}",
	GAME_130: "{userid}",
	GAME_135: "{userid}",
	GAME_136: "{userid}",
	GAME_139: "{userid}",
	GAME_140: "{rolename}|{userid}|{serverid}",
	GAME_141: "{rolename}|{userid}|{serverid}",
	GAME_142: "{rolename}|{serverid}",
	GAME_146: "{userid}",
	GAME_150: "{userid}",
	GAME_152: "{userid}",
	GAME_153: "{userid}",
	GAME_154: "{userid}|{serverid}",
	GAME_172: "{userid}|{serverid}",
	GAME_176: "{userid}|{tag}",
	GAME_230: "{userid}|{serverid}",
}

type ServerListCode string

const (
	SERVER_CODE_103 ServerListCode = "103" //Mobile Legend
	SERVER_CODE_127 ServerListCode = "127" //Ragnarok
	SERVER_CODE_140 ServerListCode = "140" //Bleach Mobile 3D
	SERVER_CODE_141 ServerListCode = "141" //Era of Celestials
	SERVER_CODE_142 ServerListCode = "142" //Dragon Nest
	SERVER_CODE_172 ServerListCode = "172" //Genshin Impact
)

type IakPostpaidType string

const (
	ASURANSI        IakPostpaidType = "asuransi"
	BPJS            IakPostpaidType = "bpjs"
	EMONEY          IakPostpaidType = "emoney"
	FINANCE         IakPostpaidType = "finance"
	GAS             IakPostpaidType = "gas"
	HP              IakPostpaidType = "hp"
	INTERNET        IakPostpaidType = "internet"
	PAJAK_DAERAH    IakPostpaidType = "pajak-daerah"
	PAJAK_KENDARAAN IakPostpaidType = "pajak-kendaraan"
	PBB             IakPostpaidType = "pbb"
	PDAM            IakPostpaidType = "pdam"
	PLN             IakPostpaidType = "pln"
	TV              IakPostpaidType = "tv"
)
