<template>
  <div class="project table-responsive">
    <h2>{{ name }}</h2>
    <table class="table table-striped table-hover table-bordered">
      <thead>
        <tr>
          <th>Service</th>
          <th>Container Name</th>
          <th>Image</th>
          <th>Domain Names</th>
          <th class="hidden-xs">Network / IP</th>
          <th class="hidden-xs">Ports</th>
          <th class="hidden-xs">Mounts</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="container in containers">
          <td>{{ container.Labels['com.docker.compose.service'] }}</td>
          <td>{{ container.Names[0] }}</td>
          <td>{{ container.Image }}</td>
          <td v-html="displayDNS(container)"></td>
          <td v-html="displayNetworks(container)" class="hidden-xs"></td>
          <td v-html="displayPorts(container)" class="hidden-xs"></td>
          <td v-html="displayRemoteMounts(container)" class="hidden-xs"></td>
        </tr>
      </tbody>
    </table>

  </div>
</template>

<script>

export default {

  name: 'Project',

  props: [ 'name', 'containers' ],

  methods: {

    displayDNS (container) {
      let names = []
      if (container.Labels.hasOwnProperty('com.dnsdock.image')) {
        names.push(container.Labels['com.dnsdock.name'] + '.' + container.Labels['com.dnsdock.image'] + '.vm')
      }
      if (container.Labels.hasOwnProperty('com.dnsdock.alias')) {
        names.push(container.Labels['com.dnsdock.alias'])
      }

      // Need to also get any VIRTUAL_HOST env entries

      let port = null
      if (container.Ports.some((port) => { return port.PrivatePort === 80 })) {
        port = 80
      } else if (container.Ports.some((port) => { return port.PrivatePort === 8080 })) {
        port = 8080
      }

      if (port) {
        names = names.map((name) => { return this.makeLink(name, port) })
      }
      return names.map(this.linebreak).join('')
    },

    displayNetworks (container) {
      let networks = []
      for (let name in container.NetworkSettings.Networks) {
        if (container.NetworkSettings.Networks.hasOwnProperty(name)) {
          let settings = container.NetworkSettings.Networks[name]
          networks.push(name + '(' + settings.IPAddress + ')')
        }
      }
      return networks.map(this.linebreak).join('')
    },

    displayPorts (container) {
      let ports = []
      container.Ports.forEach((port) => {
        let portVal = port.PrivatePort + '/' + port.Type
        ports.push(portVal)
      })
      return ports.map(this.linebreak).join('')
    },

    displayRemoteMounts (container) {
      let mounts = []
      container.Mounts.forEach((mount) => {
        if (!mount.hasOwnProperty('Driver')) {
          mounts.push(mount.Source + ':' + mount.Destination)
        }
      })
      return mounts.map(this.linebreak).join('')
    },

    linebreak (val) {
      return '<p>' + val + '</p>'
    },

    makeLink (domain, port) {
      let url = 'http://' + domain + ':' + port
      return '<a href="' + url + '" target="_blank">' + domain + '</a>'
    }
  }

}
</script>
