<template>
  <div class="home">
    <PageLoader v-if='loading'/>
    <table class="primary flex-table" v-else>
      <thead>
        <tr>
          <th>ID</th>
          <th>Type</th>
          <th>Summary</th>
          <th>Description</th>
          <th>Assignee</th>
          <th>Link</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="issue in issues" :key="issue.id">
          <td>{{ issue.id }}</td>
          <td>{{ issue.type }}</td>
          <td>{{ issue.summary }}</td>
          <td>{{ issue.description }}</td>
          <td>{{ issue.assignee }}</td>
          <td>
            <a :href="issue.link" target="_blank">
              {{ issue.link }}
            </a>
          </td>
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
      issues: []
    }
  },
  async mounted () {
    try {
      const resp = await fetch('http://localhost:8000/issues')
      const jsonResp = await resp.json()
      this.loading = false
      this.issues = jsonResp
    } catch (e) {
      console.error(e)
    }
  },
}
</script>

<style>
table {
  font-size: 12px;
}
</style>
