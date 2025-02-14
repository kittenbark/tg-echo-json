package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kittenbark/tg"
)

func main() {
	tg.NewFromEnv().
		OnError(tg.OnErrorLog).
		Filter(tg.OnPrivateMessage).
		HandleCommand("/start", tg.CommonTextReply("Hii, this bot was made with https://github.com/kittenbark/tg.")).
		Branch(tg.OnMessage, func(ctx context.Context, upd *tg.Update) error {
			msgJson, err := json.MarshalIndent(upd.Message, "", "  ")
			if err != nil {
				return err
			}

			_, err = tg.SendMessage(ctx, upd.Message.Chat.Id,
				fmt.Sprintf("```json\n%s\n```", tg.EscapeParseMode(tg.ParseModeMarkdownV2, string(msgJson))),
				&tg.OptSendMessage{ParseMode: tg.ParseModeMarkdownV2},
			)
			return err
		}).
		Start()
}
