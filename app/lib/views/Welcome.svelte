<script lang="ts">
  import Layout from "$lib/components/Layout.svelte";
  import Form from "$lib/components/forms/Form.svelte";
  import type { Form as Props } from "$lib/types/gen/main/lib/session/memory/Form";

  let { Username = "", Error = "" }: Props = $props();
</script>

<Layout title="Welcome">
  <Form method="POST" action="/">
    <input type="text" value="{Username}" name="username" class="input input-neutral" />
    <button class="btn btn-ghost">Submit</button>
  </Form>

    {#if Error}
        <div>{Error}</div>
    {/if}
</Layout>
