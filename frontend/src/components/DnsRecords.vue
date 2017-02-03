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
          <td>{{ entry.ips | join(' || ') }}</td>
        </tr>
      </tbody>
    </table>

  </div>
</template>

<script>
export default {

  name: 'DnsRecords',

  props: [ 'entries' ]
}
</script>
