Boas práticas aplicadas
1. Escolha de imagem base leve
Utilizamos uma imagem base como python:3.9-slim no Dockerfile (sugerido), que é otimizada para reduzir o tamanho da imagem final e melhorar a eficiência no uso de recursos.

2. Organização do diretório de trabalho
O comando WORKDIR /app define um diretório padrão dentro do container, garantindo que todos os comandos subsequentes sejam executados nesse contexto. Isso melhora a legibilidade e organização.

3. Instalação eficiente de dependências
As dependências são instaladas a partir do arquivo requirements.txt utilizando o comando pip install --no-cache-dir, evitando o armazenamento desnecessário de arquivos temporários e reduzindo o tamanho da imagem.

Apenas o arquivo requirements.txt é copiado inicialmente, permitindo que o cache do Docker seja aproveitado caso as dependências não mudem.

4. Separação clara de responsabilidades
O código da aplicação é copiado para o container em uma etapa separada após a instalação das dependências, seguindo o princípio de responsabilidade única para cada comando no Dockerfile.

5. Exposição explícita de portas
A porta 8000, utilizada pelo Uvicorn (servidor FastAPI), é explicitamente exposta no Dockerfile com o comando EXPOSE 8000. Isso documenta claramente qual porta será usada.

6. Comando inicial bem definido
O comando padrão para iniciar a aplicação (CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]) garante que o container seja executado corretamente assim que iniciado.

7. Uso do Docker Compose
O arquivo docker-compose.yaml facilita a execução da aplicação com:

Mapeamento de portas (8000:8000) para acessar a aplicação localmente.

Montagem de volumes (.:/app) para permitir atualizações dinâmicas no código durante o desenvolvimento.

Configuração de variáveis de ambiente (FASTAPI_ENV=development) para controlar comportamentos específicos do ambiente.

Benefícios dessas práticas
Redução do tamanho da imagem: Uso de imagens leves e eliminação de arquivos desnecessários.

Reprodutibilidade: Configuração consistente entre diferentes ambientes (desenvolvimento, teste, produção).

Eficiência: Melhor aproveitamento do cache do Docker e execução otimizada.

Manutenção facilitada: Código organizado e bem documentado, facilitando futuras alterações ou expansões.
