<template>
  <div id="dns-records" class="col">
    <h2>DNS Records</h2>
    <table class="table table-striped table-hover table-bordered"">
      <thead>
        <tr>
          <th>Domain Name</th>
          <th>IP Address</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="entry in entries">
          <td>{{ entry.name }}</td>
          <td>{{ entry.ips | join ',' }}</td>
        </tr>
      </tbody>
    </table>

  </div>
</template>

<script>
export default {

  name: 'DnsRecords',

  props: [ 'entries' ]

    /*
    {
      "580350ffd5f6b34acabc648b666905769dd072185739a5559dd534f9643bf223": {
        "Name":"dnsdock",
        "Image":"devtools",
        "IPs":["172.17.0.2"],
        "TTL":-1,
        "Aliases":[]
      },
      "a10b139072aeb626e0dc0120cf1e7a43e9d30cb4b9cde61ba43607872b95424c":{
        "Name":"web",
        "Image":"devtools",
        "IPs":["172.17.0.3"],
        "TTL":-1,
        "Aliases": [
          "other.devtools.vm"
        ]
      }
    }
    */

}
</script>
