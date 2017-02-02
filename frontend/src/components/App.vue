<template>
  <div id="app" class="container-fluid">
    <div class="jumbotron">
      <h1>Rig Dashboard</h1>
    </div>
    <div class="container">
      <div class="row">
        <dns-records></dns-records>
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
    this.dnsRecords = store.fetchDnsRecords()
  }

}
</script>

<style>
a {
  color: #42b983;
}
</style>
