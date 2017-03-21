<template>
  <div id="app" class="container">
    <div class="jumbotron">
      <div class="project-title__wrapper">
        <div class="media-left media-middle">
          <img class="logo-image" src="dist/logo.png">
        </div>
        <div class="media-body">
          <h1 class="project-name">Outrigger Dashboard</h1>
        </div>
      </div>
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

    let self = this
    this.ws = new WebSocket('ws://' + window.location.host + '/api/containers/ws')
    this.ws.addEventListener('message', function (e) {
      self.projects = store.processContainersByProject(JSON.parse(e.data))
      self.loadDnsRecords()
    })
  },

  methods: {
    update () {
      this.loadDnsRecords()
      this.loadContainers()
    },

    loadDnsRecords() {
      store.fetchDnsRecords().then(items => {
        this.dnsRecords = items
      })
    },

    loadContainers() {
      store.fetchContainersByProject().then(items => {
        this.projects = items
      })
    }
  }

}
</script>
<style>
  .jumbotron {
    background-color: #ffffff;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .logo-image {
    max-width: 150px;
  }

  .media-body {
    max-width: 650px;
  }
</style>
