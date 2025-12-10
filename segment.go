package qbot

import (
	"fmt"

	"github.com/awfufu/qbot/api"
)

// Segment represents a message segment.
type Segment api.Segment

// creates a text segment
func Text(text string) Segment {
	return Segment{
		Type: "text",
		Data: map[string]any{
			"text": text,
		},
	}
}

// creates an at segment
func At(userID uint64) Segment {
	return Segment{
		Type: "at",
		Data: map[string]any{
			"qq": fmt.Sprintf("%d", userID),
		},
	}
}

// creates a face segment
func Face(id uint64) Segment {
	return Segment{
		Type: "face",
		Data: map[string]any{
			"id": fmt.Sprintf("%d", id),
		},
	}
}

// creates an image segment
// file can be a local path (file:///path/to/file), a URL (https://example.com/image), or base64;
// summary is optional.
func Image(file string, summary ...string) Segment {
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

// Json creates a json segment
func Json(data string) Segment {
	return Segment{
		Type: "json",
		Data: map[string]any{
			"data": data,
		},
	}
}

// creates a record segment
func Record(file string) Segment {
	return Segment{
		Type: "record",
		Data: map[string]any{
			"file": file,
		},
	}
}

// creates a video segment
func Video(file string) Segment {
	return Segment{
		Type: "video",
		Data: map[string]any{
			"file": file,
		},
	}
}

// creates a music segment
func Music(typeStr, id string) Segment {
	return Segment{
		Type: "music",
		Data: map[string]any{
			"type": typeStr,
			"id":   id,
		},
	}
}

// creates a custom music segment
func CustomMusic(url, audio, title, content, image string) Segment {
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

// creates a dice segment
func Dice() Segment {
	return Segment{
		Type: "dice",
		Data: map[string]any{},
	}
}

// creates a rps segment
func Rps() Segment {
	return Segment{
		Type: "rps",
		Data: map[string]any{},
	}
}

// creates a file segment
func File(file string) Segment {
	return Segment{
		Type: "file",
		Data: map[string]any{
			"file": file,
		},
	}
}

// Node creates a forward message node using message ID.
func Node(id string) Segment {
	return Segment{
		Type: "node",
		Data: map[string]any{
			"id": id,
		},
	}
}

// CustomNode creates a custom forward message node.
func CustomNode(name string, uin uint64, content ...any) Segment {
	return Segment{
		Type: "node",
		Data: map[string]any{
			"nickname": name,
			"user_id":  fmt.Sprintf("%d", uin),
			"content":  ToMessage(content...),
		},
	}
}

// creates a reply segment
func replySegment(msgID uint64) Segment {
	return Segment{
		Type: "reply",
		Data: map[string]any{
			"id": fmt.Sprintf("%d", msgID),
		},
	}
}

// converts variadic arguments to a message ([]api.Segment)
func ToMessage(args ...any) []api.Segment {
	if len(args) == 0 {
		return []api.Segment{}
	}

	var segments []api.Segment
	for _, arg := range args {
		switch v := arg.(type) {
		case Segment:
			segments = append(segments, api.Segment(v))
		case api.Segment:
			segments = append(segments, v)
		case string:
			segments = append(segments, api.Segment(Text(v)))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			segments = append(segments, api.Segment(Text(fmt.Sprint(v))))
		case fmt.Stringer:
			segments = append(segments, api.Segment(Text(v.String())))
		default:
			// Try to convert unknown types to string representation
			segments = append(segments, api.Segment(Text(fmt.Sprintf("%v", v))))
		}
	}
	return segments
}

// ForwardBlockItem represents a single forward message node.
type ForwardBlockItem struct {
	Name    string    `json:"nickname"`
	UserID  uint64    `json:"user_id"`
	Content []Segment `json:"content"`
}

// ForwardBlock represents a block of forward messages with metadata.
type ForwardBlock struct {
	Title   string             `json:"source"`
	Preview string             `json:"preview"`
	Summary string             `json:"summary"`
	Prompt  string             `json:"prompt"`
	Content []ForwardBlockItem `json:"messages"`
}
