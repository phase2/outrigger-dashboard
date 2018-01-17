<template>
  <div id="project-list">
    <project v-for="(containers, name) in projects" :name="name" :containers="containers"></project>
  </div>
</template>

<script>
import Project from './Project.vue'

export default {

  name: 'ProjectList',

  components: {
    Project
  },

  props: [ 'projects' ]
}
</script>
