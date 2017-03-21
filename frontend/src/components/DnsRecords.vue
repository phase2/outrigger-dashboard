<template>
  <div id="dns-records" class="col">
    <h2>DNS Records</h2>
    <table class="table table-striped table-hover table-bordered">
      <thead>
        <tr>
          <th>Domain Name</th>
          <th>IP Address</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="entry in sortedEntries">
          <td><a :href="makeLink(entry.name)" target="_blank">{{ entry.name }}</a></td>
          <td>{{ entry.ips | join(' || ') }}</td>
        </tr>
      </tbody>
    </table>

  </div>
</template>

<script>
export default {

  name: 'DnsRecords',

  props: [ 'entries' ],

  methods: {
    makeLink (domainName) {
      return 'http://' + domainName
    }
  },

  computed: {
    sortedEntries() {
      // Some weirdness with this not being an array early on, so lets cast it.
      let iterable = Array.from(this.entries)
      return iterable.sort((a, b) => {
        return a.name.localeCompare(b.name)
      })
    }
  }
}
</script>
<style>
  thead {
    background-color: #0e2e41;
    color: #ffffff;
  }

  th {
    font-weight: normal;
  }

  a {
    color: #008eb0;
  }

  a:hover {
    color: #0e2e41;
    text-decoration: none;
  }

  .table-hover tbody tr:hover td {
    background-color: #E2E2E2;
  }
</style>
