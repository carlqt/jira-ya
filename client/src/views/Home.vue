<template>
  <div class="home">
    <PageLoader v-if='loading'/>
    <table class="primary" v-else>
      <thead>
        <tr>
          <th>Name</th>
          <th>City</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>Francisco</td>
          <td>Valencia, Spain</td>
        </tr>
        <tr>
          <td>Eve</td>
          <td>San Francisco, USA</td>
        </tr>
        <tr>
          <td>Martin</td>
          <td>New York, USA</td>
        </tr>
        <tr>
          <td>Sarah</td>
          <td>London, UK</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
// @ is an alias to /src
import PageLoader from '@/components/PageLoader.vue'

export default {
  name: 'Home',
  components: {
    PageLoader,
  },
  data () {
    return {
      loading: true,
    }
  },
  async mounted () {
    try {
      const resp = await fetch('http://localhost:8000/issues')
      const jsonResp = await resp.json()
      this.loading = false
      console.log(jsonResp)
    } catch (e) {
      console.error(e)
    }
  },
}
</script>
