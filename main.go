package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("토큰을 입력하세요: ")
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("start-maximized", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"),
	)
	ctx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ = chromedp.NewContext(ctx)

	loginURL := "https://discord.com/login"

	jsScript := fmt.Sprintf(`
		(function() {
			window.t = "%s";
			window.localStorage = document.body.appendChild(document.createElement('iframe')).contentWindow.localStorage;
			window.setInterval(() => window.localStorage.token = '"' + window.t + '"', 50); 
			window.location.reload();
		})();
	`, token)

	err := chromedp.Run(ctx,
		chromedp.Navigate(loginURL),
		chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(jsScript, nil),
	)
	if err != nil {
		fmt.Println("오류 발생:", err)
		return
	}

	fmt.Println("토큰이 설정되었습니다. 디스코드에 자동으로 로그인됩니다.")
}
