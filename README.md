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



