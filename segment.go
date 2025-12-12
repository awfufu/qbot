package qbot

import (
	"fmt"

	"github.com/awfufu/qbot/api"
)

type Segment = api.Segment

// create text segment
func Text(text string) Segment {
	return textSegment(text)
}

// create at segment
func At(userID uint64) Segment {
	return atSegment(userID)
}

// create face segment
func Face(id uint64) Segment {
	return faceSegment(id)
}

// create image segment
func Image(file string, summary ...string) Segment {
	return imageSegment(file, summary...)
}

func textSegment(text string) Segment {
	return Segment{
		Type: "text",
		Data: map[string]any{
			"text": text,
		},
	}
}

func atSegment(userID uint64) Segment {
	return Segment{
		Type: "at",
		Data: map[string]any{
			"qq": fmt.Sprintf("%d", userID),
		},
	}
}

func faceSegment(id uint64) Segment {
	return Segment{
		Type: "face",
		Data: map[string]any{
			"id": fmt.Sprintf("%d", id),
		},
	}
}

func imageSegment(file string, summary ...string) Segment {
	data := map[string]any{
		"file": file,
	}
	if len(summary) > 0 {
		data["summary"] = summary[0]
	}
	return Segment{
		Type: "image",
		Data: data,
	}
}

func jsonSegment(data string) Segment {
	return Segment{
		Type: "json",
		Data: map[string]any{
			"data": data,
		},
	}
}

func recordSegment(file string) Segment {
	return Segment{
		Type: "record",
		Data: map[string]any{
			"file": file,
		},
	}
}

func videoSegment(file string) Segment {
	return Segment{
		Type: "video",
		Data: map[string]any{
			"file": file,
		},
	}
}

func musicSegment(typeStr, id string) Segment {
	return Segment{
		Type: "music",
		Data: map[string]any{
			"type": typeStr,
			"id":   id,
		},
	}
}

func customMusicSegment(url, audio, title, content, image string) Segment {
	return Segment{
		Type: "music",
		Data: map[string]any{
			"type":    "custom",
			"url":     url,
			"audio":   audio,
			"title":   title,
			"content": content,
			"image":   image,
		},
	}
}

func diceSegment() Segment {
	return Segment{
		Type: "dice",
		Data: map[string]any{},
	}
}

func rpsSegment() Segment {
	return Segment{
		Type: "rps",
		Data: map[string]any{},
	}
}

func fileSegment(file string) Segment {
	return Segment{
		Type: "file",
		Data: map[string]any{
			"file": file,
		},
	}
}

func nodeSegment(id string) Segment {
	return Segment{
		Type: "node",
		Data: map[string]any{
			"id": id,
		},
	}
}

func customNodeSegment(name string, uin uint64, content ...any) Segment {
	return Segment{
		Type: "node",
		Data: map[string]any{
			"nickname": name,
			"user_id":  fmt.Sprintf("%d", uin),
			"content":  ToMessage(content...),
		},
	}
}

func replySegment(msgID uint64) Segment {
	return Segment{
		Type: "reply",
		Data: map[string]any{
			"id": fmt.Sprintf("%d", msgID),
		},
	}
}

func ToMessage(args ...any) []api.Segment {
	if len(args) == 0 {
		return []api.Segment{}
	}

	var segments []api.Segment
	for _, arg := range args {
		switch v := arg.(type) {
		case Segment:
			segments = append(segments, v)
		case string:
			segments = append(segments, textSegment(v))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			segments = append(segments, textSegment(fmt.Sprint(v)))
		case fmt.Stringer:
			segments = append(segments, textSegment(v.String()))
		default:
			// Try to convert unknown types to string representation
			segments = append(segments, textSegment(fmt.Sprintf("%v", v)))
		}
	}
	return segments
}

// represents a block of forward messages with metadata.
type ForwardBlock struct {
	Title   string             `json:"source"`
	Preview string             `json:"preview"`
	Summary string             `json:"summary"`
	Prompt  string             `json:"prompt"`
	Content []ForwardBlockItem `json:"messages"`
}

// represents a single forward message node.
type ForwardBlockItem struct {
	Name    string    `json:"nickname"`
	UserID  uint64    `json:"user_id"`
	Content []Segment `json:"content"`
}
