
module.exports = {
    devServer: {
       port: 8080, // CHANGE YOUR PORT HERE!
       proxy: {
           '/api': {
                   target: 'http://localhost:3000', // provide proxy  for your project
                   ws: true,
                   changeOrigin: true,
                   pathRewrite: {
                       '^/api': ''
               }
           },
           '/cdn': {
               target: 'http://localhost:8887/', // provide proxy  for your project
               ws: true,
               changeOrigin: true,
               pathRewrite: {
                   '^/cdn': ''
           }
       }
       }
   }
   }