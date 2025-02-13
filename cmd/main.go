package main

import (
	"context"
	"encoding/json"
	"github.com/kittenbark/tg"
	"unicode/utf8"
)

func main() {
	tg.NewFromEnv().
		OnError(tg.OnErrorLog).
		Filter(tg.OnPrivateMessage).
		HandleCommand("/start", tg.CommonTextReply("Hii, this bot was made with https://github.com/kittenbark/tg.")).
		Branch(tg.OnMessage, func(ctx context.Context, upd *tg.Update) error {
			data, err := json.MarshalIndent(upd.Message, "", "  ")
			if err != nil {
				return err
			}

			text := string(data)
			_, err = tg.SendMessage(ctx, upd.Message.Chat.Id, text, &tg.OptSendMessage{
				Entities: []*tg.MessageEntity{{
					Type:     "pre",
					Offset:   0,
					Length:   int64(utf8.RuneCountInString(text)) + 2,
					Language: "json",
				}}})
			return err
		}).
		Start()
}
