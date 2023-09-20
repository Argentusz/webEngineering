<template>
<b-modal
  id="university-edit-modal"
  header-bg-variant="warning"
  header-text-variant="dark"
  header-border-variant="warning"
  body-bg-variant="dark"
  body-text-variant="white"
  footer-bg-variant="dark"
  footer-text-variant="white"
  footer-border-variant="warning"
  cancel-disabled
  ok-disabled
>
  <template #modal-footer>
    <b-button
        class="text-sm"
        @click="$bvModal.hide('university-edit-modal'); $emit('hide')"
    >
      {{ $t('cancel') }}
    </b-button>
    <b-button
        variant="success"
        class="text-sm"
        @click="saveHandler"
    >
      {{ $t('save') }}
    </b-button>
  </template>

</b-modal>
</template>

<script>
export default {
  name: "UniversityEdit"
}
</script>

<style scoped>

</style>