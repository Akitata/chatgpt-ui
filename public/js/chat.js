
var bot = new ChatSDK({

    config: {
        navbar: {
            title: 'ChatGPT Bot'
        },
        messages: [
            {
                type: 'text',
                content: {
                    text: 'ChatGPT Bot为您服务，请问有什么可以帮您？'
                }
            }
        ]
    },
    requests: {
        send: function (msg) {
            if (msg.type === 'text') {
                return {
                    url: '/chat',
                    data: {
                        prompt: msg.content.text
                    }
                };
            }
        },
    },
});

bot.run();
