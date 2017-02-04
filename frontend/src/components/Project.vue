<template>
  <div class="col project">
    <h2>{{ name }}</h2>
    <table class="table table-striped table-hover table-bordered"">
      <thead>
        <tr>
          <th>Service</th>
          <th>Container Name</th>
          <th>Image</th>
          <th>Network / IP</th>
          <th>Ports</th>
          <th>Mounts</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="container in containers">
          <td>{{ container.Labels['com.docker.compose.service'] }}</td>
          <td>{{ container.Names[0] }}</td>
          <td>{{ container.Image }}</td>
          <td v-html="displayNetworks(container)"></td>
          <td v-html="displayPorts(container)"></td>
          <td v-html="displayRemoteMounts(container)"></td>
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

    displayNetworks(container) {
      let networks = [];
      for (let name in container.NetworkSettings.Networks) {
        if ( container.NetworkSettings.Networks.hasOwnProperty(name) ) {
          let settings = container.NetworkSettings.Networks[name];
          networks.push(name + "(" + settings.IPAddress + ")")
        }
      }
      return networks.map(this.linebreak).join('')
    },

    displayPorts(container) {
      let ports = [];
      container.Ports.forEach( (port) => {
        let portVal = port.PrivatePort + "/" + port.Type;

        if ([80, 443, 8080].includes(port.PrivatePort)) {
          portVal = this.makeLink(container, portVal, port.PrivatePort == "443")
        }

        ports.push(portVal)
      });
      return ports.map(this.linebreak).join('')
    },

    displayRemoteMounts(container) {
      let mounts = [];
      container.Mounts.forEach( (mount) => {
        if ( !mount.hasOwnProperty("Driver") ) {
          mounts.push(mount.Source + ":" + mount.Destination)
        }
      });
      return mounts.map(this.linebreak).join('')
    },

    linebreak (val) {
      return "<p>" + val + "</p>";
    },

    makeLink (container, text, secure) {
      if ( container.Labels.hasOwnProperty('com.dnsdock.image') ) {
        let url = secure ? "https://" : "http://"
        url += container.Labels['com.dnsdock.name'] + "." + container.Labels['com.dnsdock.image'] + ".vm"
        return "<a href='" + url + "' target='_blank'>" + text + "</a>"
      }
      return text
    }
  }

}
</script>
