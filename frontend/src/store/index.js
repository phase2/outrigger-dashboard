import { EventEmitter } from 'events'
import 'whatwg-fetch'

const store = new EventEmitter();

export default store

/**
 * Fetch dns records.
 */
store.fetchDnsRecords = () => {
  return fetch('/api/dnsrecords')
    .then(function(response) {
      return response.json()
    })
    .then(function(json) {
      return store.processDnsRecords(json)
    })
    .catch(function(ex) {
      console.log(ex);
      return {};
  })
};

/**
 * Take all the dnsdock records and generate the DNS names
 */
store.processDnsRecords = (records) => {
  let entries = [];
  for (let prop in records) {
    if( records.hasOwnProperty(prop) ) {
      let record = records[prop];

      entries.push({
        ips: record.IPs,
        name: record.Name + '.' + record.Image + '.vm'
      });

      record.Aliases.forEach(alias => {
        entries.push({
          ips: record.IPs,
          name: alias
        });
      });
    }
  }
  return entries;
};


/**
 * Get all running containers
 */
store.fetchContainers = () => {
  return fetch('/api/containers')
    .then(function(response) {
      return response.json();
    })
    .catch(function(ex) {
      console.log(ex);
      return {};
    })
};

/**
 * Get all containers and organize them by compose project name
 */
store.fetchContainersByProject = () => {
  return store.fetchContainers()
    .then(function(json) {
      return store.processContainersByProject(json);
    })
    .catch(function(ex) {
      console.log(ex);
      return {};
    })
};

/**
 * Take an array of all containers and organize them by Compose project name
 */
store.processContainersByProject = (containers) => {
  let projects = {};
  containers.forEach((container) => {
    let projectName = "Unclassified";
    if (container.Labels.hasOwnProperty('com.docker.compose.project')) {
      projectName = container.Labels['com.docker.compose.project'];
    }

    if (!projects.hasOwnProperty(projectName)) {
      projects[projectName] = [];
    }

    projects[projectName].push(container);
  });
  return projects;
};

