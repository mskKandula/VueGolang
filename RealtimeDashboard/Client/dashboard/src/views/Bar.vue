<template>
  <div>
    <h3>Bar Chart Example in Vue</h3>
    <bar-chart :data12 = this.data12></bar-chart>
  </div>
</template>

<script>
import BarChart from '@/components/BarChart'

export default {
  components: {
    BarChart
  },
   data() {
      return {
     data12 :[]
      }
   },
   created() {
  var loc = window.location;
        var uri = 'ws:';

        if (loc.protocol === 'https:') {
          uri = 'wss:';
        }
        uri += '//' + loc.host;
        uri += loc.pathname + 'ws';
      
        this.ws = new WebSocket(uri)

        this.ws.onconnect = (evt) => {
          console.log("ws connected", evt)
        }
        this.ws.onmessage = (evt) =>{

          let x = JSON.parse(evt.data)
         this.data12 = x.Intarr
       
        }
        this.ws.onclose = (evt) => {
          console.log("ws closed", evt)
        }
  }
 
}
</script>