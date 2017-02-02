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
    let records = store.fetchDnsRecords()
    for (var prop in records) {
      if( records.hasOwnProperty(prop) ) {
        Vue.set(this.dnsRecords, prop, records[prop])
      }
    }
  }

}
</script>

<style>
a {
  color: #42b983;
}
</style>
