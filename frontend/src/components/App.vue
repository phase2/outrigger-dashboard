<template>
  <div id="app" class="container-fluid">
    <div class="jumbotron">
      <h1>Outrigger Dashboard</h1>
    </div>
    <div class="container">
      <div class="row">
        <dns-records :entries="dnsRecords"></dns-records>
      </div>
      <project-list :projects="projects"></project-list>
    </div>
  </div>
</template>

<script>
import DnsRecords from './DnsRecords.vue'
import ProjectList from './ProjectList.vue'
import store from '../store'


export default {
  name: 'app',
  data () {
    return {
      dnsRecords: {},
      projects: {}
    }
  },
  components: {
    DnsRecords,
    ProjectList
  },

  created () {
    this.update()
    store.on('dataloaded', this.update)
  },

  methods: {
    update() {
      store.fetchDnsRecords().then(items => {
        this.dnsRecords = items;
      })
      store.fetchContainersByProject().then(items => {
        this.projects = items;
      })
    }
  }

}
</script>
