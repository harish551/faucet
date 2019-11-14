<template lang="pug">
.section
  .section-header
    h1 Community Links
  .section-main
    list-item(
      href="https://e-money.network/"
      target="_blank"
      title="Explorer"
      subtitle="#explorer")
    list-item(
      href="https://t.me/emoney_com"
      target="_blank"
      title="Telegram"
      subtitle="#emoney_com")
    list-item(
      href="https://github.com/e-money"
      target="_blank"
      title="Source Code"
      subtitle="e-money")
    list-item(
      href="https://e-money.com/"
      target="_blank"
      title="Website"
      subtitle="e-money.com")
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
