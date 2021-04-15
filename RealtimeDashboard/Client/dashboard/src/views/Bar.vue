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
   }
 
}
</script>