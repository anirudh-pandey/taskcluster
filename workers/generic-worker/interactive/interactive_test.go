package interactive

import (
	"context"
	"testing"

	"github.com/gorilla/websocket"
)

func TestInteractive(t *testing.T) {
	// Start an interactive session on a test server
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	interactive, err := New(53654, ctx)
	if err != nil {
		t.Fatalf("could not create interactive session: %v", err)
	}
	done := make(chan error, 1)
	go func() {
		done <- interactive.ListenAndServe()
	}()

	// Make a WebSocket connection to the server
	url := "ws://localhost:53654"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatal("dial error:", err)
	}

	// Send some input to the interactive session
	input := "echo hello\n"
	err = conn.WriteMessage(websocket.TextMessage, []byte(input))
	if err != nil {
		t.Fatal("write error:", err)
	}

	// Wait for the output from the interactive session
	_, output, err := conn.ReadMessage()
	if err != nil {
		t.Fatal("read error:", err)
	}
	expected := "hello\n"
	if string(output) != expected {
		t.Fatalf("unexpected output: %v\nexpected: %v", string(output), expected)
	}

	input = "notABashCommand\n"
	err = conn.WriteMessage(websocket.TextMessage, []byte(input))
	if err != nil {
		t.Fatal("write error:", err)
	}

	// Wait for the output from the interactive session
	_, output, err = conn.ReadMessage()
	if err != nil {
		t.Fatal("read error:", err)
	}
	expected = "bash: line 2: notABashCommand: command not found\n"
	if string(output) != expected {
		t.Fatalf("unexpected output: %v\nexpected: %v", string(output), expected)
	}

	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Closing connection"))
	if err != nil {
		t.Fatalf("Error sending WebSocket close message: %v", err)
	}

	_, _, err = conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
			t.Fatalf("Unexpected close error: %v", err)
		}
	}

	err = conn.Close()
	if err != nil {
		t.Fatalf("Error closing WebSocket connection: %v", err)
	}

	// Terminate the interactive session
	err = interactive.Terminate()
	if err != nil {
		t.Fatal("terminate error:", err)
	}

	err = <-done
	if err != nil {
		t.Fatalf("listen and serve error: %v", err)
	}
}
