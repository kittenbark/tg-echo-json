package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kittenbark/tg"
)

func main() {
	tg.NewFromEnv().
		Scheduler(). // Optionally you could add Scheduler() to ensure no 429.
		OnError(tg.OnErrorLog). // Let's log errors with slog.
		Branch(tg.OnMessage, func(ctx context.Context, upd *tg.Update) error { // If tg update is message, send its contents.
			msg := upd.Message
			msgJson, err := json.MarshalIndent(msg, "", "  ")
			if err != nil {
				return err
			}

			_, err = tg.SendMessage(ctx, msg.Chat.Id,
				fmt.Sprintf("```json\n%s\n```", tg.EscapeParseMode(tg.ParseModeMarkdownV2, string(msgJson))),
				&tg.OptSendMessage{
					ParseMode:       tg.ParseModeMarkdownV2,
					ReplyParameters: &tg.ReplyParameters{MessageId: msg.MessageId},
				},
			)
			return err
		}).
		Filter(tg.OnPrivateMessage). // Filter everything below, allow only private messages.
		Command("/start", tg.CommonTextReply(helloMessage)). // Send greeting message.
		Start()
}

const helloMessage = "Hii, this bot was made with https://github.com/kittenbark/tg."
