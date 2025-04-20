package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Uso: converter <input.mp4> <output.mp4>")
        os.Exit(1)
    }

    input := os.Args[1]
    output := os.Args[2]

    log.Printf("Convertendo %s para %s...\n", input, output)

    // Chama o ffmpeg para converter o vídeo
    cmd := exec.Command("ffmpeg", "-y", "-i", input, "-c:v", "libx264", "-c:a", "aac", output)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        log.Fatalf("Erro ao converter vídeo: %v", err)
    }

    log.Println("Conversão concluída!")
}
