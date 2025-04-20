# desafio-docker-images

## 01 - Primeiro Docker File
  
### build
![01](https://github.com/user-attachments/assets/3e851bdd-9082-46b3-abd1-c18d18546b2f)

### deploy
![02](https://github.com/user-attachments/assets/7b2665de-6658-4491-a4a6-bf8e36d43ba5)

### App Online
![03](https://github.com/user-attachments/assets/4f71e195-6990-4228-a737-a54675e3ec08)


## 02 - Boas Práticas
### build / deploy
![00](https://github.com/user-attachments/assets/eb0b983f-0fb0-4e66-ac64-464171d9e961)

### App Online
![01](https://github.com/user-attachments/assets/148fb829-76b9-4d69-ab9d-4c492946f0fe)


## 03 - Entrypoint VS CMD
### Aplicação para conversão de Vídeos Utilizando Entrypoint e CMD: chamando a função de conversão de vídeos com Docker Compose, seguida processamento da tarefa e resultado.
docker-compose run video-converter example.MP4 avi
Iniciando conversão de 'example.MP4' para 'output.avi'...
ffmpeg version 4.2.7-0ubuntu0.1 Copyright (c) 2000-2022 the FFmpeg developers
  built with gcc 9 (Ubuntu 9.4.0-1ubuntu1~20.04.1)
  configuration: --prefix=/usr --extra-version=0ubuntu0.1 --toolchain=hardened --libdir=/usr/lib/x86_64-linux-gnu --incdir=/usr/include/x86_64-linux-gnu --arch=amd64 --enable-gpl --disable-stripping --enable-avresample --disable-filter=resample --enable-avisynth --enable-gnutls --enable-ladspa --enable-libaom --enable-libass --enable-libbluray --enable-libbs2b --enable-libcaca --enable-libcdio --enable-libcodec2 --enable-libflite --enable-libfontconfig --enable-libfreetype --enable-libfribidi --enable-libgme --enable-libgsm --enable-libjack --enable-libmp3lame --enable-libmysofa --enable-libopenjpeg --enable-libopenmpt --enable-libopus --enable-libpulse --enable-librsvg --enable-librubberband --enable-libshine --enable-libsnappy --enable-libsoxr --enable-libspeex --enable-libssh --enable-libtheora --enable-libtwolame --enable-libvidstab --enable-libvorbis --enable-libvpx --enable-libwavpack --enable-libwebp --enable-libx265 --enable-libxml2 --enable-libxvid --enable-libzmq --enable-libzvbi --enable-lv2 --enable-omx --enable-openal --enable-opencl --enable-opengl --enable-sdl2 --enable-libdc1394 --enable-libdrm --enable-libiec61883 --enable-nvenc --enable-chromaprint --enable-frei0r --enable-libx264 --enable-shared
  libavutil      56. 31.100 / 56. 31.100
  libavcodec     58. 54.100 / 58. 54.100
  libavformat    58. 29.100 / 58. 29.100
  libavdevice    58.  8.100 / 58.  8.100
  libavfilter     7. 57.100 /  7. 57.100
  libavresample   4.  0.  0 /  4.  0.  0
  libswscale      5.  5.100 /  5.  5.100
  libswresample   3.  5.100 /  3.  5.100
  libpostproc    55.  5.100 / 55.  5.100
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from 'example.MP4':
  Metadata:
    major_brand     : isom
    minor_version   : 512
    compatible_brands: isomiso2avc1mp41
    encoder         : Lavf60.3.100
  Duration: 00:03:44.21, start: 0.000000, bitrate: 445 kb/s
    Stream #0:0(und): Video: h264 (High) (avc1 / 0x31637661), yuv420p(tv, bt709), 640x360 [SAR 1:1 DAR 16:9], 306 kb/s, 30 fps, 30 tbr, 15360 tbn, 60 tbc (default)
    Metadata:
      handler_name    : ISO Media file produced by Google Inc.
      encoder         : Lavc60.3.100 libx264
    Stream #0:1(eng): Audio: aac (LC) (mp4a / 0x6134706D), 44100 Hz, stereo, fltp, 130 kb/s (default)
    Metadata:
      handler_name    : ISO Media file produced by Google Inc.
Stream mapping:
  Stream #0:0 -> #0:0 (h264 (native) -> h264 (libx264))
  Stream #0:1 -> #0:1 (aac (native) -> aac (native))
Press [q] to stop, [?] for help
[libx264 @ 0x555d26396100] using SAR=1/1
[libx264 @ 0x555d26396100] using cpu capabilities: MMX2 SSE2Fast SSSE3 SSE4.2 AVX FMA3 BMI2 AVX2
[libx264 @ 0x555d26396100] profile High, level 3.0
Output #0, avi, to 'output.avi':
  Metadata:
    major_brand     : isom
    minor_version   : 512
    compatible_brands: isomiso2avc1mp41
    ISFT            : Lavf58.29.100
    Stream #0:0(und): Video: h264 (libx264) (H264 / 0x34363248), yuv420p, 640x360 [SAR 1:1 DAR 16:9], q=-1--1, 30 fps, 30 tbn, 30 tbc (default)
    Metadata:
      handler_name    : ISO Media file produced by Google Inc.
      encoder         : Lavc58.54.100 libx264
    Side data:
      cpb: bitrate max/min/avg: 0/0/0 buffer size: 0 vbv_delay: -1
    Stream #0:1(eng): Audio: aac (LC) ([255][0][0][0] / 0x00FF), 44100 Hz, stereo, fltp, 128 kb/s (default)
    Metadata:
      handler_name    : ISO Media file produced by Google Inc.
      encoder         : Lavc58.54.100 aac
frame=  453 fps=0.0 q=29.0 size=     522kB time=00:00:15.55 bitrate= 274.7kbiframe=  958 fps=958 q=29.0 size=    1546kB time=00:00:32.39 bitrate= 390.9kbiframe= 1483 fps=988 q=29.0 size=    2314kB time=00:00:49.89 bitrate= 379.8kbiframe= 2029 fps=1014 q=29.0 size=    3338kB time=00:01:08.08 bitrate= 401.6kbframe= 2564 fps=1025 q=29.0 size=    4106kB time=00:01:25.91 bitrate= 391.5kbframe= 3102 fps=1033 q=29.0 size=    5130kB time=00:01:43.86 bitrate= 404.6kbframe= 3567 fps=1018 q=29.0 size=    5898kB time=00:01:59.35 bitrate= 404.8kbframe= 4063 fps=1015 q=29.0 size=    6922kB time=00:02:15.85 bitrate= 417.4kbframe= 4581 fps=1016 q=29.0 size=    7690kB time=00:02:33.15 bitrate= 411.3kbframe= 5102 fps=1019 q=29.0 size=    8714kB time=00:02:50.52 bitrate= 418.6kbframe= 5612 fps=1019 q=29.0 size=    9482kB time=00:03:07.50 bitrate= 414.3kbframe= 6130 fps=1020 q=29.0 size=   10506kB time=00:03:24.77 bitrate= 420.3kbframe= 6643 fps=1021 q=29.0 size=   11274kB time=00:03:41.88 bitrate= 416.2kbframe= 6725 fps=1015 q=-1.0 Lsize=   11770kB time=00:03:44.21 bitrate= 430.0kbits/s speed=33.9x    
video:7852kB audio:3516kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 3.535460%
[libx264 @ 0x555d26396100] frame I:27    Avg QP:17.82  size: 32896
[libx264 @ 0x555d26396100] frame P:1827  Avg QP:21.47  size:  2705
[libx264 @ 0x555d26396100] frame B:4871  Avg QP:27.29  size:   454
[libx264 @ 0x555d26396100] consecutive B-frames:  1.0%  6.8%  1.9% 90.3%
[libx264 @ 0x555d26396100] mb I  I16..4: 13.1% 35.1% 51.7%
[libx264 @ 0x555d26396100] mb P  I16..4:  1.9%  3.1%  1.2%  P16..4: 20.1%  7.8%  4.5%  0.0%  0.0%    skip:61.4%
[libx264 @ 0x555d26396100] mb B  I16..4:  0.2%  0.3%  0.1%  B16..8: 11.6%  2.4%  0.5%  direct: 1.0%  skip:83.9%  L0:42.7% L1:51.0% BI: 6.4%
[libx264 @ 0x555d26396100] 8x8 transform intra:48.5% inter:57.7%
[libx264 @ 0x555d26396100] coded y,uvDC,uvAC intra: 48.7% 60.2% 26.0% inter: 4.1% 4.5% 0.5%
[libx264 @ 0x555d26396100] i16 v,h,dc,p: 33% 24% 15% 28%
[libx264 @ 0x555d26396100] i8 v,h,dc,ddl,ddr,vr,hd,vl,hu: 29% 17% 26%  4%  4%  5%  4%  6%  5%
[libx264 @ 0x555d26396100] i4 v,h,dc,ddl,ddr,vr,hd,vl,hu: 29% 18% 14%  6%  7%  7%  7%  7%  6%
[libx264 @ 0x555d26396100] i8c dc,h,v,p: 47% 19% 26%  7%
[libx264 @ 0x555d26396100] Weighted P-Frames: Y:0.4% UV:0.3%
[libx264 @ 0x555d26396100] ref P L0: 77.7% 22.3%
[libx264 @ 0x555d26396100] ref B L0: 87.7% 12.3%
[libx264 @ 0x555d26396100] ref B L1: 96.0%  4.0%
[libx264 @ 0x555d26396100] kb/s:286.95
[aac @ 0x555d263a53c0] Qavg: 241.359
Conversão concluída com sucesso! Arquivo gerado: output.avi

## 04 - Multistage Build

docker compose up --build

Compose can now delegate builds to bake for better performance.
 To do so, set COMPOSE_BAKE=true.
[+] Building 16.4s (19/19) FINISHED                                                          docker:default
 => [bbb-converter internal] load build definition from Dockerfile                                     0.0s
 => => transferring dockerfile: 1.25kB                                                                 0.0s
 => [bbb-converter internal] load metadata for docker.io/library/alpine:3.21                          11.6s
 => [bbb-converter internal] load metadata for docker.io/library/golang:1.24-alpine3.21               11.6s
 => [bbb-converter internal] load .dockerignore                                                        0.0s
 => => transferring context: 2B                                                                        0.0s
 => [bbb-converter builder 1/7] FROM docker.io/library/golang:1.24-alpine3.21@sha256:7772cb5322baa875  0.0s
 => [bbb-converter internal] load build context                                                        0.0s
 => => transferring context: 1.78kB                                                                    0.0s
 => [bbb-converter stage-1 1/5] FROM docker.io/library/alpine:3.21@sha256:a8560b36e8b8210634f77d9f7f9  0.0s
 => CACHED [bbb-converter builder 2/7] RUN apk update && apk add --no-cache git ca-certificates tzdat  0.0s
 => CACHED [bbb-converter builder 3/7] RUN mkdir -p /srv/bbb-convert /srv/bbb-convert/release          0.0s
 => CACHED [bbb-converter builder 4/7] WORKDIR /srv/bbb-convert                                        0.0s
 => [bbb-converter builder 5/7] COPY . .                                                               0.1s
 => [bbb-converter builder 6/7] RUN go mod download                                                    1.0s
 => [bbb-converter builder 7/7] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -e  3.6s
 => CACHED [bbb-converter stage-1 2/5] RUN apk add --no-cache chromium ffmpeg                          0.0s
 => CACHED [bbb-converter stage-1 3/5] RUN adduser -D bigbluebutton bigbluebutton                      0.0s
 => CACHED [bbb-converter stage-1 4/5] COPY --from=builder /srv/bbb-convert/release /srv/bbb-convert   0.0s
 => CACHED [bbb-converter stage-1 5/5] WORKDIR /home/bigbluebutton                                     0.0s
 => [bbb-converter] exporting to image                                                                 0.0s
 => => exporting layers                                                                                0.0s
 => => writing image sha256:d95249311ffeece8dc8155be6526067c03be5e757bc35794a217233a8b4732d5           0.0s
 => => naming to docker.io/library/04-multistage-build-bbb-converter                                   0.0s
 => [bbb-converter] resolving provenance for metadata file                                             0.0s
[+] Running 2/2
 ✔ bbb-converter                                  Built                                                0.0s 
 ✔ Container 04-multistage-build-bbb-converter-1  C...                                                 0.0s 
Attaching to bbb-converter-1
bbb-converter-1  | 2025/04/20 16:42:47 Convertendo /videos/input.mp4 para /videos/output.mp4...
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
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] using SAR=1/1
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] using cpu capabilities: MMX2 SSE2Fast SSSE3 SSE4.2 AVX FMA3 BMI2 AVX2
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] profile High, level 4.2, 4:2:0, 8-bit
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] 264 - core 164 r3108 - H.264/MPEG-4 AVC codec - Copyleft 2003-2023 - http://www.videolan.org/x264.html - options: cabac=1 ref=3 deblock=1:0:0 analyse=0x3:0x113 me=hex subme=7 psy=1 psy_rd=1.00:0.00 mixed_ref=1 me_range=16 chroma_me=1 trellis=1 8x8dct=1 cqm=0 deadzone=21,11 fast_pskip=1 chroma_qp_offset=-2 threads=24 lookahead_threads=4 sliced_threads=0 nr=0 decimate=1 interlaced=0 bluray_compat=0 constrained_intra=0 bframes=3 b_pyramid=2 b_adapt=1 b_bias=0 direct=1 weightb=1 open_gop=0 weightp=2 keyint=250 keyint_min=25 scenecut=40 intra_refresh=0 rc_lookahead=40 rc=crf mbtree=1 crf=23.0 qcomp=0.60 qpmin=0 qpmax=69 qpstep=4 ip_ratio=1.40 aq=1:1.00
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
frame=   53 fps= 49 q=31.0 size=       0kB time=00:00:00.85 bitrate=   0.5kbits/s dup=63 drop=0 speed=0.794xframe=  125 fps= 79 q=31.0 size=     256kB time=00:00:02.03 bitrate=1030.5kbits/s dup=98 drop=0 speed=1.29x frame=  199 fps= 96 q=31.0 size=     256kB time=00:00:03.26 bitrate= 641.5kbits/s dup=135 drop=0 speed=1.57xframe=  266 fps=103 q=31.0 size=     256kB time=00:00:04.40 bitrate= 476.2kbits/s dup=169 drop=0 speed=1.71xframe=  341 fps=110 q=31.0 size=     256kB time=00:00:05.65 bitrate= 370.9kbits/s dup=207 drop=0 speed=1.83xframe=  418 fps=116 q=31.0 size=     256kB time=00:00:06.94 bitrate= 302.2kbits/s dup=245 drop=0 speed=1.93xframe=  497 fps=121 q=31.0 size=     512kB time=00:00:08.25 bitrate= 507.9kbits/s dup=284 drop=0 speed=2.01xframe=  581 fps=126 q=31.0 size=     512kB time=00:00:09.65 bitrate= 434.2kbits/s dup=326 drop=0 speed=2.09xframe=  661 fps=129 q=31.0 size=     512kB time=00:00:10.99 bitrate= 381.5kbits/s dup=368 drop=0 speed=2.15xframe=  736 fps=131 q=31.0 size=     512kB time=00:00:12.22 bitrate= 343.0kbits/s dup=405 drop=0 speed=2.18xframe=  826 fps=135 q=31.0 size=     512kB time=00:00:13.74 bitrate= 305.1kbits/s dup=453 drop=0 speed=2.24xframe=  913 fps=138 q=31.0 size=     768kB time=00:00:15.19 bitrate= 414.0kbits/s dup=496 drop=0 speed=2.29xframe= 1000 fps=140 q=31.0 size=     768kB time=00:00:16.63 bitrate= 378.3kbits/s dup=539 drop=0 speed=2.33xframe= 1086 fps=142 q=31.0 size=     768kB time=00:00:18.08 bitrate= 347.9kbits/s dup=584 drop=0 speed=2.37xframe= 1177 fps=145 q=31.0 size=    1024kB time=00:00:19.60 bitrate= 427.9kbits/s dup=629 drop=0 speed=2.41xframe= 1265 fps=146 q=31.0 size=    1024kB time=00:00:21.07 bitrate= 398.1kbits/s dup=673 drop=0 speed=2.44xframe= 1335 fps=146 q=31.0 size=    1280kB time=00:00:22.23 bitrate= 471.5kbits/s dup=708 drop=0 speed=2.43xframe= 1417 fps=147 q=31.0 size=    1536kB time=00:00:23.59 bitrate= 533.4kbits/s dup=749 drop=0 speed=2.44xframe= 1504 fps=148 q=31.0 size=    1536kB time=00:00:25.05 bitrate= 502.2kbits/s dup=793 drop=0 speed=2.46xframe= 1578 fps=148 q=31.0 size=    1536kB time=00:00:26.29 bitrate= 478.6kbits/s dup=829 drop=0 speed=2.46xframe= 1634 fps=146 q=31.0 size=    1536kB time=00:00:27.22 bitrate= 462.2kbits/s dup=857 drop=0 speed=2.44xframe= 1698 fps=145 q=31.0 size=    1792kB time=00:00:28.29 bitrate= 518.8kbits/s dup=890 drop=0 speed=2.42xframe= 1774 fps=146 q=31.0 size=    1792kB time=00:00:29.54 bitrate= 496.9kbits/s dup=927 drop=0 speed=2.42xframe= 1850 fps=146 q=31.0 size=    2048kB time=00:00:30.81 bitrate= 544.5kbits/s dup=965 drop=0 speed=2.43x[vost#0:0/libx264 @ 0x7fc945153380] More than 1000 frames duplicated
[out#0/mp4 @ 0x7fc945280680] video:7594kB audio:0kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 1.181348%    
bbb-converter-1  | frame= 7617 fps=157 q=-1.0 Lsize=    7684kB time=00:02:07.02 bitrate= 495.5kbits/s dup=3835 drop=0 speed=2.62x    
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] frame I:32    Avg QP:17.99  size: 80635
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] frame P:2000  Avg QP:23.51  size:  2012
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] frame B:5585  Avg QP:35.40  size:   210
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] consecutive B-frames:  0.9%  3.1%  2.7% 93.3%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] mb I  I16..4: 43.3% 36.9% 19.8%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] mb P  I16..4:  1.6%  1.1%  0.4%  P16..4:  1.1%  0.2%  0.1%  0.0%  0.0%    skip:95.5%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] mb B  I16..4:  0.1%  0.0%  0.0%  B16..8:  1.9%  0.0%  0.0%  direct: 0.0%  skip:98.0%  L0:40.6% L1:59.2% BI: 0.2%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] 8x8 transform intra:35.1% inter:48.5%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] coded y,uvDC,uvAC intra: 12.8% 9.4% 5.1% inter: 0.1% 0.1% 0.0%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] i16 v,h,dc,p: 21% 78%  0%  0%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] i8 v,h,dc,ddl,ddr,vr,hd,vl,hu: 41% 15% 40%  0%  0%  0%  1%  0%  2%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] i4 v,h,dc,ddl,ddr,vr,hd,vl,hu: 28% 33% 16%  2%  4%  4%  4%  3%  5%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] i8c dc,h,v,p: 71% 26%  3%  0%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] Weighted P-Frames: Y:0.1% UV:0.1%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] ref P L0: 70.5%  4.4% 17.3%  7.8%  0.0%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] ref B L0: 70.5% 26.9%  2.6%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] ref B L1: 97.4%  2.6%
bbb-converter-1  | [libx264 @ 0x7fc93fb48940] kb/s:489.52
bbb-converter-1  | 2025/04/20 16:43:30 Conversão concluída!
bbb-converter-1 exited with code 0

