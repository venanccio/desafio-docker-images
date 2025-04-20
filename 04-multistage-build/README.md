### Desafio 04: Multistage Build
## Projeto: Conversor de Vídeo com Go, Docker e FFmpeg 

## Descrição
Este projeto é um conversor de vídeo simples desenvolvido em Go, que utiliza FFmpeg para processar arquivos de vídeo. O ambiente é totalmente containerizado com Docker, usando uma build multistage para gerar imagens leves e seguras.
O objetivo é aceitar um arquivo de vídeo de entrada (ex: input.mp4) e gerar um arquivo convertido de saída (ex: output.mp4), facilitando testes, automação e integração em pipelines.

## Como usar
1. Pré-requisitos
Docker

## Docker Compose

2. Preparando o ambiente
Coloque o vídeo de entrada na pasta videos com o nome input.mp4:

text
./videos/input.mp4
3. Comando para build e execução
No terminal, execute:

text
docker compose up --build
Exemplo de execução
Abaixo está um exemplo real do log de execução do projeto, mostrando o processo de build da imagem, inicialização do container e conversão do vídeo:

text
docker compose up --build
Compose can now delegate builds to bake for better performance.
 To do so, set COMPOSE_BAKE=true.
[+] Building 25.5s (19/19) FINISHED                                                                    docker:default
 => [bbb-converter internal] load build definition from Dockerfile                                               0.0s
 => => transferring dockerfile: 1.25kB                                                                           0.0s
 => [bbb-converter internal] load metadata for docker.io/library/alpine:3.21                                    20.7s
 => [bbb-converter internal] load metadata for docker.io/library/golang:1.24-alpine3.21                         20.7s
 => [bbb-converter internal] load .dockerignore                                                                  0.0s
 => => transferring context: 2B                                                                                  0.0s
 => [bbb-converter builder 1/7] FROM docker.io/library/golang:1.24-alpine3.21@sha256:7772cb5322baa875edd7470555  0.0s
 => [bbb-converter stage-1 1/5] FROM docker.io/library/alpine:3.21@sha256:a8560b36e8b8210634f77d9f7f9efd7ffa463  0.0s
 => [bbb-converter internal] load build context                                                                  0.0s
 => => transferring context: 12.75kB                                                                             0.0s
 => CACHED [bbb-converter builder 2/7] RUN apk update && apk add --no-cache git ca-certificates tzdata && updat  0.0s
 => CACHED [bbb-converter builder 3/7] RUN mkdir -p /srv/bbb-convert /srv/bbb-convert/release                    0.0s
 => CACHED [bbb-converter builder 4/7] WORKDIR /srv/bbb-convert                                                  0.0s
 => [bbb-converter builder 5/7] COPY . .                                                                         0.2s
 => [bbb-converter builder 6/7] RUN go mod download                                                              1.0s
 => [bbb-converter builder 7/7] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags   3.5s
 => CACHED [bbb-converter stage-1 2/5] RUN apk add --no-cache chromium ffmpeg                                    0.0s
 => CACHED [bbb-converter stage-1 3/5] RUN adduser -D bigbluebutton bigbluebutton                                0.0s
 => CACHED [bbb-converter stage-1 4/5] COPY --from=builder /srv/bbb-convert/release /srv/bbb-convert             0.0s
 => CACHED [bbb-converter stage-1 5/5] WORKDIR /home/bigbluebutton                                               0.0s
 => [bbb-converter] exporting to image                                                                           0.0s
 => => exporting layers                                                                                          0.0s
 => => writing image sha256:d95249311ffeece8dc8155be6526067c03be5e757bc35794a217233a8b4732d5                     0.0s
 => => naming to docker.io/library/04-multistage-build-bbb-converter                                             0.0s
 => [bbb-converter] resolving provenance for metadata file                                                       0.0s
[+] Running 2/2
 ✔ bbb-converter                                  Built                                                          0.0s 
 ✔ Container 04-multistage-build-bbb-converter-1  Created                                                        0.0s 
Attaching to bbb-converter-1
bbb-converter-1  | 2025/04/20 16:48:27 Convertendo /videos/input.mp4 para /videos/output.mp4...
bbb-converter-1  | ffmpeg version 6.1.2 Copyright (c) 2000-2024 the FFmpeg developers
bbb-converter-1  |   built with gcc 14.2.0 (Alpine 14.2.0)
bbb-converter-1  |   configuration: --prefix=/usr --disable-librtmp --disable-lzma --disable-static --disable-stripping --enable-avfilter --enable-gpl --enable-ladspa --enable-libaom --enable-libass --enable-libbluray --enable-libdav1d --enable-libdrm --enable-libfontconfig --enable-libfreetype --enable-libfribidi --enable-libharfbuzz --enable-libmp3lame --enable-libopenmpt --enable-libopus --enable-libplacebo --enable-libpulse --enable-librav1e --enable-librist --enable-libsoxr --enable-libsrt --enable-libssh --enable-libtheora --enable-libv4l2 --enable-libvidstab --enable-libvorbis --enable-libvpx --enable-libwebp --enable-libx264 --enable-libx265 --enable-libxcb --enable-libxml2 --enable-libxvid --enable-libzimg --enable-libzmq --enable-lto=auto --enable-lv2 --enable-openssl --enable-pic --enable-postproc --enable-pthreads --enable-shared --enable-vaapi --enable-vdpau --enable-version3 --enable-vulkan --optflags=-O3 --enable-libjxl --enable-libsvtav1 --enable-libvpl
bbb-converter-1  |   libavutil      58. 29.100 / 58. 29.100
bbb-converter-1  |   libavcodec     60. 31.102 / 60. 31.102
bbb-converter-1  |   libavformat    60. 16.100 / 60. 16.100
bbb-converter-1  |   libavdevice    60.  3.100 / 60.  3.100
bbb-converter-1  |   libavfilter     9. 12.100 /  9. 12.100
bbb-converter-1  |   libswscale      7.  5.100 /  7.  5.100
bbb-converter-1  |   libswresample   4. 12.100 /  4. 12.100
bbb-converter-1  |   libpostproc    57.  3.100 / 57.  3.100
bbb-converter-1  | Input #0, mov,mp4,m4a,3gp,3g2,mj2, from '/videos/input.mp4':
bbb-converter-1  |   Metadata:
bbb-converter-1  |     major_brand     : isom
bbb-converter-1  |     minor_version   : 512
bbb-converter-1  |     compatible_brands: isomiso2avc1mp41
bbb-converter-1  |     encoder         : Lavf58.76.100
bbb-converter-1  |   Duration: 00:02:07.08, start: 0.000000, bitrate: 1318 kb/s
bbb-converter-1  |   Stream #0:0[0x1](eng): Video: h264 (Baseline) (avc1 / 0x31637661), yuv420p(tv, bt709, progressive), 1920x1080, 1315 kb/s, SAR 1:1 DAR 16:9, 29.77 fps, 59.94 tbr, 16k tbn (default)
bbb-converter-1  |     Metadata:
bbb-converter-1  |       handler_name    : VideoHandler
bbb-converter-1  |       vendor_id       : [0][0][0][0]
bbb-converter-1  | Stream mapping:
bbb-converter-1  |   Stream #0:0 -> #0:0 (h264 (native) -> h264 (libx264))
bbb-converter-1  | Press [q] to stop, [?] for help
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] using SAR=1/1
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] using cpu capabilities: MMX2 SSE2Fast SSSE3 SSE4.2 AVX FMA3 BMI2 AVX2
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] profile High, level 4.2, 4:2:0, 8-bit... bbb-converter-1  | [libx264 @ 0x7ff5796a8940] 264 - core 164 r3108 - H.264/MPEG-4 AVC codec - Copyleft 2003-2023 - http://www.videolan.org/x264.html - options: cabac=1 ref=3 deblock=1:0:0 analyse=0x3:0x113 me=hex subme=7 psy=1 psy_rd=1.00:0.00 mixed_ref=1 me_range=16 chroma_me=1 trellis=1 8x8dct=1 cqm=0 deadzone=21,11 fast_pskip=1 chroma_qp_offset=-2 threads=24 lookahead_threads=4 sliced_threads=0 nr=0 decimate=1 interlaced=0 bluray_compat=0 constrained_intra=0 bframes=3 b_pyramid=2 b_adapt=1 b_bias=0 direct=1 weightb=1 open_gop=0 weightp=2 keyint=250 keyint_min=25 scenecut=40 intra_refresh=0 rc_lookahead=40 rc=crf mbtree=1 crf=23.0 qcomp=0.60 qpmin=0 qpmax=69 qpstep=4 ip_ratio=1.40 aq=1:1.00
bbb-converter-1  | Output #0, mp4, to '/videos/output.mp4':
bbb-converter-1  |   Metadata:
bbb-converter-1  |     major_brand     : isom
bbb-converter-1  |     minor_version   : 512
bbb-converter-1  |     compatible_brands: isomiso2avc1mp41
bbb-converter-1  |     encoder         : Lavf60.16.100
bbb-converter-1  |   Stream #0:0(eng): Video: h264 (avc1 / 0x31637661), yuv420p(tv, bt709, progressive), 1920x1080 [SAR 1:1 DAR 16:9], q=2-31, 59.94 fps, 60k tbn (default)
bbb-converter-1  |     Metadata:
bbb-converter-1  |       handler_name    : VideoHandler
bbb-converter-1  |       vendor_id       : [0][0][0][0]
bbb-converter-1  |       encoder         : Lavc60.31.102 libx264
bbb-converter-1  |     Side data:
bbb-converter-1  |       cpb: bitrate max/min/avg: 0/0/0 buffer size: 0 vbv_delay: N/A
[vost#0:0/libx264 @ 0x7ff57ec95a80] More than 1000 frames duplicated 606.3kbits/s dup=974 drop=0 speed=2.55x     
[out#0/mp4 @ 0x7ff57ecb3680] video:7594kB audio:0kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 1.181348%    
bbb-converter-1  | frame= 7617 fps=164 q=-1.0 Lsize=    7684kB time=00:02:07.02 bitrate= 495.5kbits/s dup=3835 drop=0 speed=2.73x    
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] frame I:32    Avg QP:17.99  size: 80635
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] frame P:2000  Avg QP:23.51  size:  2012
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] frame B:5585  Avg QP:35.40  size:   210
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] consecutive B-frames:  0.9%  3.1%  2.7% 93.3%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] mb I  I16..4: 43.3% 36.9% 19.8%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] mb P  I16..4:  1.6%  1.1%  0.4%  P16..4:  1.1%  0.2%  0.1%  0.0%  0.0%    skip:95.5%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] mb B  I16..4:  0.1%  0.0%  0.0%  B16..8:  1.9%  0.0%  0.0%  direct: 0.0%  skip:98.0%  L0:40.6% L1:59.2% BI: 0.2%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] 8x8 transform intra:35.1% inter:48.5%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] coded y,uvDC,uvAC intra: 12.8% 9.4% 5.1% inter: 0.1% 0.1% 0.0%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] i16 v,h,dc,p: 21% 78%  0%  0%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] i8 v,h,dc,ddl,ddr,vr,hd,vl,hu: 41% 15% 40%  0%  0%  0%  1%  0%  2%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] i4 v,h,dc,ddl,ddr,vr,hd,vl,hu: 28% 33% 16%  2%  4%  4%  4%  3%  5%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] i8c dc,h,v,p: 71% 26%  3%  0%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] Weighted P-Frames: Y:0.1% UV:0.1%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] ref P L0: 70.5%  4.4% 17.3%  7.8%  0.0%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] ref B L0: 70.5% 26.9%  2.6%
bbb-converter-1  | [libx264 @ 0x7ff5796a8940] ref B L1: 97.4%  2.6%
bbb-converter-1  | kb/s:489.52
bbb-converter-1  | 2025/04/20 16:49:11 Conversão concluída!
bbb-converter-1 exited with code 0

## Estrutura de pastas esperada
text
.
├── Dockerfile
├── docker-compose.yaml
├── main.go
├── go.mod
├── go.sum
└── videos/
    ├── input.mp4
    └── output.mp4  # gerado após a conversão

## Observações
O projeto utiliza uma imagem multistage para garantir builds rápidos e imagens enxutas.

O binário Go chama o FFmpeg para realizar a conversão do vídeo.

O volume ./videos é mapeado para o container, facilitando a entrada e saída dos arquivos.

## Licença
Distribuído sob a licença MIT. Sinta-se à vontade para usar, modificar e contribuir!

Se precisar de mais instruções ou exemplos, consulte os comentários no código ou abra uma issue.
