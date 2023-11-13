<script lang="ts">
	let show = false
</script>
	<!-- Id string `json:"id"` -->
	<!-- Name     string `json:"name"` -->
	<!-- Roaster  string `json:"roaster"` -->
	<!-- Country  string `json:"country"` -->
	<!-- Varietal string `json:"varietal"` -->
	<!-- Process  string `json:"process"` -->
	<!-- Altitude string `json:"altitude"` -->
	<!-- Notes    string `json:"notes"` -->
	<!-- Weight   float32 `json:"weight"` -->

{#if show}
	<div class="card">
		<form method="POST">
			<label class="label">Name: <input class="input" name="name"/> </label>
			<label class="label">Roaster: <input class="input" name="roaster"/> </label>
			<label class="label">Weight: <input class="input" name="weight"/> </label>
			<label class="label">Country: <input class="input" name="country"/> </label>
			<label class="label">Varietal <input class="input" name="varietal"/> </label>
			<label class="label">Altitude <input class="input" name="altitude"/> </label>
			<label class="label">Notes: <input class="input" name="notes"/> </label>
		</form>
	</div>
{/if}
<button on:click={() => show = true}>+ Add Bean</button>

