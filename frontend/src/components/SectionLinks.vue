<template lang="pug">
.section
  .section-header
    h1 Community Links
  .section-main
    list-item(
      href="https://vitwit.com/"
      target="_blank"
      title="Website"
      subtitle="#vitwit.com")
    list-item(
      href="https://twitter.com/vitwit_india"
      target="_blank"
      title="Twitter"
      subtitle="#vitwit_india")
    list-item(
      href="https://github.com/vitwit"
      target="_blank"
      title="Github"
      subtitle="#vitwit")
    list-item(
      href="https://medium.com/vitwit"
      target="_blank"
      title="Medium"
      subtitle="#vitwit")
</template>

<script>
import ListItem from "../components/NiListItem.vue";
export default {
  name: "faucet",
  components: {
    ListItem
  },
  data: () => ({}),
  methods: {}
};
</script>
