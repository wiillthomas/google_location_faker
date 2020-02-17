
<script>
	import { fade } from "svelte/transition"
	
	let location;
	let query;
	let googleLink;
	let flashError;
	let locationPlaceholder = "Fr";

	function handleClick(){
		if ( location && query ) {
			fetch(`http://localhost:8080/api?q=${ query }&location=${ location }` )
			.then( ( response ) =>  response.json() )
			.then( ( data ) => {
				googleLink = `https://google.com/search?${data}`
				flashError = null;
			} )
			.catch( err => {
				flashError = "An issue occoured while communicating with the API. Email help@thisdomain.com if the problem persists."
			} )
		} else {
			flashError = "Please fill in the required fields"
		}
	}

	function handleAutoComplete() {
		console.log(location)
		fetch(`http://localhost:8080/autocomplete?input=${location}`)
		.then(( response ) => {
			return response.json()
		})
		.then((data) => {
			console.log(data)  
			locationPlaceholder = ""
			for ( let i = 0; i < 10; i += 1 ) {
				if ( data[ i ] !== undefined ) {
					locationPlaceholder += `${location}${data[i]}
					`
				}
			}
		})
		.catch(err =>  console.log(err))
	}

</script>

<style>
	h1 {
		color: purple;
	}
	.flash-message {
		display: flex;
		height: 50px;
		background: #bb0000;
		color: white;
		position: absolute;
		top: 0px;
		opacity: 1;
		left: 0px;
		width: 100vw;
		justify-content: center;
		align-items: center;
	}
	.location-input__container {
		position: relative;
	}
	.location-input__container input {
		position: absolute;  
		font-size: 16px; 
		letter-spacing: 1px;  
	}
	.location-input__container .placeholder {
		position: absolute;
		font-size: 14px;
		top: 38px;
		left: 20px;
    	letter-spacing: 1px; 
	}
</style>

{#if flashError}
	<div class="flash-message">
		{ flashError }
	</div>
{/if}

	<input bind:value={query}>
<div class="location-input__container">
	<input bind:value={location} on:keyup={handleAutoComplete} placeholder={locationPlaceholder}>
	<div class="placeholder">{ locationPlaceholder }</div>
</div>

<button on:click={handleClick}>
	Get Results
</button>

{#if googleLink}
	<a href={googleLink}>link here!</a>
{/if}