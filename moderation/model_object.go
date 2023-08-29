package moderation

// 开发者文本检测响应
type TextModerationOutput struct {
	Type  string
	Score float64
}

// 文本同步审核请求
type TextSyncModerationInput struct {
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

// ===================================中间数据结构===================================
type Text struct {
	ContentId string `json:"contentId"` // 唯一id
	Data      string `json:"data"`      // 文本内容
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
