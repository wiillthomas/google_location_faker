
<script>
	import { fade } from "svelte/transition"
	
	let location;
	let query;
	let googleLink;
	let flashError;
	let locationPlaceholder = "Fr";
	let possibleLocations = []
	let possibleLocationFocus = -1;
	let autoCompleteError = false;

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

	function handlePossibleLocationSubmit( e ) {
		e.preventDefault() 
		location = e.currentTarget.name
	}

	function handleLocationInputChange( e ) {
		if ( e.keyCode == 38 ) {
			if ( possibleLocationFocus !== -1 ){
				e.preventDefault()
				possibleLocationFocus -= 1
				document.getElementById(`possibleLocation--${ possibleLocationFocus }`).focus()
			}
		} else if ( e.keyCode == 40 ) {
			e.preventDefault()
			possibleLocationFocus += 1
			document.getElementById(`possibleLocation--${ possibleLocationFocus }`).focus()
		} else {
			possibleLocations.length = 0
			possibleLocationFocus = -1
			fetch(`http://localhost:8080/autocomplete?input=${location}`)
			.then(( response ) => {
				autoCompleteError = false
				return response.json()
			})
			.then((data) => {
				locationPlaceholder = ""
				possibleLocations.length = 0
				for ( let i = 0; i < 10; i += 1 ) {
					if ( data[i] ){
						possibleLocations.push(`${location}${data[i]}`)
					}
				}
	
			})
			.catch(( err ) =>  {
				if ( location ) {
					autoCompleteError = true
				}
			})
		}
	}

</script>

<style>
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

	.container {
		height: 100vh;
		width: 100vw;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	
	.container .input-container {
		display: flex;
		flex-direction: row;
	}
	.container input, .container button {
		height: 40px;
	}

	input.query {
		position: relative;
	}
	.possible-location-container {
		position: absolute;
		display: flex;
		flex-direction: column
	}

	.possible-location-container button {
		text-align: left;
	}


	.autoCompleteError {
		position: relative;
	}

	.error-tooltip {
		display : none;
		position: absolute;
		background: black;
		padding: 15px 20px;
    	top: 55px;
		left: 5px;
		z-index: 500;
		visibility: visible;
		color: white;
	}

	.error-tooltip::after {
		left: 10px;
		content: "";
		background-color: inherit;
		height: 13px;
		width: 13px;
		position: absolute;
		top: -6px;
		transform: rotate(45deg);
	}
	
	.autoCompleteError .error-tooltip {
		display: block;
	}

</style>

{#if flashError}
	<div class="flash-message">
		{ flashError }
	</div>
{/if}

<div class="container">
	<div class="input-container">
		<input class="query" bind:value={query}>
		<div class="location location-input__container { autoCompleteError ? "autoCompleteError" : "" }">
			<input bind:value={location} on:keyup={handleLocationInputChange} >
				<div class="error-tooltip">
					No location found in our database - however you can still try it. Note the input is case sensitive.
				</div>
				<div class="possible-location-container">
					{#each possibleLocations as location, i}
						<button on:click={handlePossibleLocationSubmit} on:keyup={handleLocationInputChange} id={`possibleLocation--${i}`} name={location}>
							{location} 
						</button>
					{/each}
				</div>
		</div>
		<button on:click={handleClick}>
			Get Results
		</button>
	</div>
	<div class="link-container"> 
		{#if googleLink}
			<a href={googleLink}>link here!</a>
		{/if}
	</div>
</div>