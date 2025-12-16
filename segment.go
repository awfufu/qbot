package qbot

import (
	"fmt"
	"log"

	"github.com/awfufu/qbot/api"
)

type Segment = api.Segment

// create text segment
func Text(text string) Segment {
	return textSegment(text)
}

// create at segment
func At(userID int64) Segment {
	return atSegment(userID)
}

// create face segment
func Face(id int16) Segment {
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

func atSegment(userID int64) Segment {
	var target string
	if userID == int64(AtAll) {
		target = "all"
	} else {
		target = fmt.Sprintf("%d", userID)
	}
	return Segment{
		Type: "at",
		Data: map[string]any{
			"qq": target,
		},
	}
}

func faceSegment(id int16) Segment {
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

// func nodeSegment(id string) Segment {
// 	return Segment{
// 		Type: "node",
// 		Data: map[string]any{
// 			"id": id,
// 		},
// 	}
// }

func customNodeSegment(name string, uin uint64, content ...any) Segment {
	return Segment{
		Type: "node",
		Data: map[string]any{
			"nickname": name,
			"user_id":  fmt.Sprintf("%d", uin),
			"content":  toSegments(content...),
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

func rawArrayToSegments(array []MsgItem) []Segment {
	var segments []Segment
	for _, e := range array {
		switch v := e.(type) {
		case TextItem:
			segments = append(segments, textSegment(v.String()))
		case AtItem:
			segments = append(segments, atSegment(int64(v)))
		case FaceItem:
			segments = append(segments, faceSegment(int16(v)))
		case *ImageItem:
			segments = append(segments, imageSegment(v.Url))
		}
	}
	return segments
}

func toSegments(args ...any) []Segment {
	if len(args) == 0 {
		return []Segment{}
	}

	var segments []Segment
	for _, arg := range args {
		switch v := arg.(type) {
		case TextItem:
			if v == "" {
				log.Println("[WARN] continue empty text segment")
				continue
			}
			segments = append(segments, textSegment(v.String()))
		case AtItem:
			if int64(v) <= 0 {
				log.Printf("[WARN] continue invalid [at:%d] segment", v)
				continue
			}
			segments = append(segments, atSegment(int64(v)))
		case FaceItem:
			if int16(v) < 0 {
				log.Printf("[WARN] continue invalid [face:%d] segment", v)
				continue
			}
			segments = append(segments, faceSegment(int16(v)))
		case *ImageItem:
			if v == nil {
				log.Println("[WARN] continue nil image segment")
				continue
			}
			segments = append(segments, imageSegment(v.Url))
		case []MsgItem:
			if len(v) == 0 {
				log.Println("[WARN] continue empty msg item array")
				continue
			}
			segments = append(segments, rawArrayToSegments(v)...)
		case Segment:
			segments = append(segments, v)
		case string:
			if v == "" {
				log.Println("[WARN] continue empty string segment")
				continue
			}
			segments = append(segments, textSegment(v))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			segments = append(segments, textSegment(fmt.Sprint(v)))
		case fmt.Stringer:
			if v == nil {
				log.Println("[WARN] continue nil stringer segment")
				continue
			}
			s := v.String()
			if s == "" {
				log.Println("[WARN] continue empty stringer segment")
				continue
			}
			segments = append(segments, textSegment(s))
		default:
			// Try to convert unknown types to string representation
			if v == nil {
				log.Println("[WARN] continue nil segment")
				continue
			}
			s := fmt.Sprintf("%v", v)
			if s == "" {
				log.Println("[WARN] continue empty segment")
				continue
			}
			segments = append(segments, textSegment(s))
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
