<template>
  <div id="app" class="container-fluid">
    <div class="jumbotron">
      <h1>Rig Dashboard</h1>
    </div>
    <div class="container">
      <div class="row">
        <dns-records :entries="dnsRecords"></dns-records>
        <!--<websites></websites>-->
      </div>
    </div>
  </div>
</template>

<script>
import DnsRecords from './DnsRecords.vue'
import store from '../store'


export default {
  name: 'app',
  data () {
    return {
      dnsRecords: {}
    }
  },
  components: {
    DnsRecords
  },

  created () {
    console.log("App.created")
    this.update()
    store.on('dataloaded', this.update)
  },

  methods: {
    update() {
      console.log("App.update")
      store.fetchDnsRecords().then(items => {
        console.log("Fetched items")
        console.log(items)
        this.dnsRecords = items;
      })
    }
  }

}
</script>

<style>
a {
  color: #42b983;
}
</style>
