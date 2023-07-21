package main

import (
	"log"
	"os"
	"os/exec"
)

func runGoFile(filePath string) {
	cmd := exec.Command("go", "run", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// List of paths to the three Go files
	goFiles := []string{
		"C:/Users/ASUS/Desktop/GO-lang/Encryption/AES/sms_producer/smsproducer.go",
		"C:/Users/ASUS/Desktop/GO-lang/Encryption/AES/Email_producer/emailproducer.go",
		"C:/Users/ASUS/Desktop/GO-lang/Encryption/AES/Whatsapp_producer/whatsappproducer.go",
	}

	// Run each Go file in sequence
	for _, goFile := range goFiles {
		runGoFile(goFile)
	}

	log.Println("All processes finished.")
}
