package moderation

// 开发者文本检测响应
type TextModerationOutput struct {
	Type  string
	Score float64
}

// 文本审核请求
type TextModerationInput struct {
	Text       []Text `json:"text"`       // 文本请求内容
	Extra      string `json:"extra"`      // 额外信息，会在响应/回调中返回
	BusinessId string `json:"businessId"` // 文本业务id
}

// 文本同步审核响应
type TextSyncModerationOutput struct {
	ModerationType   int                    `json:"moderationType"`   // 审核类型枚举
	ModerationResult []TextModerationResult `json:"moderationResult"` // 审核结果
	RequestId        string                 `json:"requestId"`        // 当次审核请求id
	Extra            string                 `json:"extra"`            // 透传字段
}

// 图片审核请求
type ImageModerationInput struct {
	Images     []Image `json:"images"`     // 图片请求内容
	Extra      string  `json:"extra"`      // 额外信息，会在响应/回调中返回
	BusinessId string  `json:"businessId"` // 图片业务id
}

// 图片同步审核响应
type ImageSyncModerationOutput struct {
	ModerationType   int                     `json:"moderationType"`   // 审核类型枚举
	ModerationResult []ImageModerationResult `json:"moderationResult"` // 审核结果
	RequestId        string                  `json:"requestId"`        // 当次审核请求id
	Extra            string                  `json:"extra"`            // 透传字段
}

// ===================================中间数据结构===================================
type Text struct {
	ContentId string `json:"contentId"` // 唯一id
	Data      string `json:"data"`      // 文本内容
}

type Image struct {
	ContentId string `json:"contentId"` // 唯一id
	Data      string `json:"data"`      // 图片链接/base64
	Type      int    `json:"type"`      // 图片类型，1:图片链接,url  2:base64
}

// 文本审核结果
type TextModerationResult struct {
	Code              int           `json:"code"`              // 状态码。200:成功，408:文本长度超限，430:业务不可用
	Msg               string        `json:"message"`           // 状态码对应的消息
	MachineSuggestion int           `json:"machineSuggestion"` // 机审建议:1正常，2疑似违规，3违规
	MachineTagL1      string        `json:"machineTagL1"`      // 机审一级标签，当machineSuggestion为1时，返回”正常“
	MachineTagL2      string        `json:"machineTagL2"`      // 机审二级标签，当machineSuggestion为1时，为空
	ContentId         string        `json:"contentId"`         // 审核时传入的contentId
	UniqueId          string        `json:"uniqueId"`          // 该条数据对应的请求标识,由ScanA生成
	MatchedList       []MatchedList `json:"matchedList"`       // 文本命中的关键词列表
	AllTags           []TextAllTag  `json:"allTags"`           // 命中的所有违规标签以及其详情信息
}

// 所有违规标签信息
type TextAllTag struct {
	MachineSuggestion int           `json:"machineSuggestion"` // 机审建议:1正常，2疑似违规，3违规
	MachineTagL1      string        `json:"machineTagL1"`      // 一级标签
	MachineTagL2      string        `json:"machineTagL2"`      // 二级标签
	Confidence        float64       `json:"confidence"`        // 置信度，值在0到1之间，值越大，违规可能性越高。
	MatchedList       []MatchedList `json:"matchedList"`       // 该段文本命中关键词时存在，字段详情参考外层matchedList
}

// 关键词匹配信息
type MatchedList struct {
	Keyword     string `json:"keyword"`     // 命中的关键词
	Tag         string `json:"tag"`         // 该敏感词的一级二级标签拼接，如“涉政:领导人”
	Description string `json:"description"` // 关键词其他描述
	Position    []int  `json:"position"`    // 该敏感词在此段文本中的启止位置。如[3,5]，代表该文段的第3个字符和第4个字符命中了关键词
}

// 图片审核结果
type ImageModerationResult struct {
	Code              int                       `json:"code"`              // 状态码。200:成功，408:文本长度超限，430:业务不可用
	Msg               string                    `json:"message"`           // 状态码对应的消息
	MachineSuggestion int                       `json:"machineSuggestion"` // 机审建议:1正常，2疑似违规，3违规
	MachineTagL1      string                    `json:"machineTagL1"`      // 机审一级标签，当machineSuggestion为1时，返回”正常“
	MachineTagL2      string                    `json:"machineTagL2"`      // 机审二级标签，当machineSuggestion为1时，为空
	ContentId         string                    `json:"contentId"`         // 审核时客户传入的contentId
	UniqueId          string                    `json:"uniqueId"`          // 该条数据对应的请求标识,由ScanA生成
	RiskDetails       ImageModerationRiskDetail `json:"riskDetails"`       // 违规风险详情
	AllTags           []ImageAllTag             `json:"allTags"`           // 命中的所有违规标签以及其详情信息
}

// 图片违规风险详情
type ImageModerationRiskDetail struct {
	Faces      []ImageModerationRiskDetailFace   `json:"faces"`      // 返图片中敏感人物的名称及位置信息
	OcrText    []ImageModerationRiskDetailOcr    `json:"ocrText"`    // 图片中违规文字相关信息
	Objects    []ImageModerationRiskDetailObject `json:"objects"`    // 物体信息，返回图片中违规物体信息
	RiskSource int                               `json:"riskSource"` // 标识违规点来源，1: 无风险，2: 文字风险，3: 图片视觉风险
}

// 图片所有违规标签信息
// 违规可能会命中多种违规，最终违规信息参考外层
type ImageAllTag struct {
	MachineSuggestion int                       `json:"machineSuggestion"` // 审建议:1正常，2疑似违规，3违规
	MachineTagL1      string                    `json:"machineTagL1"`      // 机审一级标签，当machineSuggestion为1时，返回”正常“
	MachineTagL2      string                    `json:"machineTagL2"`      // 机审二级标签，当machineSuggestion为1时，为空
	Confidence        float64                   `json:"confidence"`        // 置信度，0到1之间，值越高，可信度越高
	Details           ImageModerationRiskDetail `json:"details"`           // 当前违规类别的风险详情
}

type ImageModerationRiskDetailObject struct {
	Id         int       `json:"id"`         // 该物体在此图上的id编号
	Location   []float64 `json:"location"`   // 物体框坐标。该数组有四个值，分别代表目标框左上角和右下角的坐标。如[33,44,55,66]<br/>33代表的是左上角的x坐标,44代表左上角的y坐标,55代表的是右下角的x坐标,66代表的是右下角的y坐标
	Name       string    `json:"name"`       // 违规标签
	Confidence float64   `json:"confidence"` // 置信度，0到1之间，值越高，可信度越高
}

type ImageModerationRiskDetailFace struct {
	Id         int       `json:"id"`         // 该人脸在此图上的id编号，如果同一个人脸在图片中出现n次，则会有多个id
	Location   []float64 `json:"location"`   // 人脸框坐标。该数组有四个值，分别代表目标框左上角和右下角的坐标。如[33,44,55,66]33代表的是左上角的x坐标,44代表左上角的y坐标,55代表的是右下角的x坐标,66代表的是右下角的y坐标
	Name       string    `json:"name"`       // 人物名称
	Ratio      float64   `json:"ratio"`      // 人脸大小在整张图片的占比，0到1之间，值越高，说明人脸越大，更加值得关注
	Confidence float64   `json:"confidence"` // 置信度，0到1之间，值越高，可信度越高
}

type ImageModerationRiskDetailOcr struct {
	Location      []float64                  `json:"location"`      // 文本框坐标。该数组有八个值，分别代表文本框四个顶点的x和y坐标
	Text          string                     `json:"text"`          // 该段文本
	MatchedList   []MatchedList              `json:"matchList"`     // 当该段文本命中关键词时存在
	ModelRiskList []ImageModerationModelRisk `json:"modelRiskList"` // 当该段文本命中语义模型时存在
}

type ImageModerationModelRisk struct {
	Description string  `json:"description"` // 命中语义模型类别结果，如“涉政:敏感事件”（一级二级标签拼接）
	Confidence  float64 `json:"confidence"`  // 置信度，0到1之间，值越高，可信度越高
}
