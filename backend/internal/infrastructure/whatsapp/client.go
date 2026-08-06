package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
)

type Client struct {
	Token         string
	PhoneNumberID string
}

func NewClient(token, phoneNumberID string) *Client {
	return &Client{
		Token:         token,
		PhoneNumberID: phoneNumberID,
	}
}

// 1. Enviar mensaje de texto normal
func (c *Client) Send(phone, message string) error {
	reqBody := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                phone,
		"type":              "text",
		"text": map[string]string{
			"body": message,
		},
	}

	body, _ := json.Marshal(reqBody)
	url := fmt.Sprintf("https://graph.facebook.com/v23.0/%s/messages", c.PhoneNumberID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBytes, _ := io.ReadAll(resp.Body)
		log.Printf("[META RESP]: %s", string(respBytes))
		return fmt.Errorf("error whatsapp status %d: %s", resp.StatusCode, string(respBytes))
	}
	return nil
}

// 2. Subir el PDF a Meta (Corregido con Content-Type: application/pdf)
func (c *Client) UploadDocument(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error abriendo archivo: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("messaging_product", "whatsapp")

	// --- INICIO DE LA CORRECCIÓN ---
	// Forzamos el Content-Type a application/pdf para que Meta lo acepte
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filepath.Base(filePath)))
	h.Set("Content-Type", "application/pdf")

	part, err := writer.CreatePart(h)
	if err != nil {
		return "", err
	}
	// --- FIN DE LA CORRECCIÓN ---

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://graph.facebook.com/v23.0/%s/media", c.PhoneNumberID)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("error uploading media: %s", string(respBody))
	}

	var res struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(respBody, &res); err != nil {
		return "", err
	}

	return res.ID, nil
}

// 3. Enviar el PDF adjunto
func (c *Client) SendDocument(phone, mediaID, filename string) error {
	reqBody := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                phone,
		"type":              "document",
		"document": map[string]string{
			"id":       mediaID,
			"filename": filename,
		},
	}

	body, _ := json.Marshal(reqBody)
	url := fmt.Sprintf("https://graph.facebook.com/v23.0/%s/messages", c.PhoneNumberID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBytes, _ := io.ReadAll(resp.Body)
		log.Printf("[META RESP DOC]: %s", string(respBytes))
		return fmt.Errorf("error enviando documento %d: %s", resp.StatusCode, string(respBytes))
	}

	return nil
}
