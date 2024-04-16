# go-websocket
Simple implemetation of websocket using golang language

Handshake inicial: O processo de comunicação WebSocket começa com um handshake HTTP entre o cliente e o servidor. O cliente envia uma solicitação HTTP especial chamada de "WebSocket handshake request" para o servidor, solicitando a atualização da conexão para WebSocket. O servidor, se estiver configurado para suportar WebSockets, responde com uma "WebSocket handshake response", indicando que está pronto para estabelecer a conexão WebSocket.

Estabelecimento da conexão: Após o handshake, a conexão WebSocket é estabelecida entre o cliente e o servidor. Esta conexão é mantida aberta e permite que ambos os lados enviem dados um para o outro a qualquer momento, sem a necessidade de uma nova solicitação HTTP.

Comunicação bidirecional: Uma vez estabelecida a conexão, tanto o cliente quanto o servidor podem enviar mensagens um para o outro a qualquer momento, sem depender de solicitações HTTP. Isso permite uma comunicação bidirecional em tempo real entre as partes.

Mensagens: As mensagens enviadas através da conexão WebSocket podem ser de dois tipos: texto ou binário. As mensagens de texto são mensagens codificadas em UTF-8, enquanto as mensagens binárias são usadas para dados binários brutos. Ambos os tipos de mensagens podem ser enviados e recebidos de forma assíncrona.

Baixa sobrecarga: Uma das vantagens dos WebSockets é que eles têm uma sobrecarga menor em comparação com outras formas de comunicação em tempo real, como o uso repetido de solicitações HTTP ou o uso de técnicas como polling ou long-polling.

Encerramento da conexão: A conexão WebSocket pode ser encerrada pelo cliente ou pelo servidor a qualquer momento. Quando isso acontece, ambas as partes podem enviar uma mensagem de encerramento para notificar a outra parte sobre o encerramento da conexão.