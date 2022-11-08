<script>


let fen = "";

const regex = /^((([rnbkqpPRNBKQ1-8]{1,8})\/?){8})\s([wbWB]{1})\s([-KQkq]{1,4})\s[1-8a-hA-H-]{1,2}\s[01]{1}\s(?!0)[0-9]{1,2}$/g
let test = regex.test("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")


async function getBoard(fen){
    var formdata = new FormData();
formdata.append("Fen", fen);

var requestOptions = {
  method: 'POST',
  body: formdata,
 
};

try{
    let resp = await fetch("/api/adv", requestOptions)
    return resp.text();
} catch (error){
    console.log(error)
}

}

let promise = getBoard()

const validFen = () => {

    
   
    return "Valid Fen!"}



</script>
<div class="flex w-full justify-center bg-teal-100">
    <div class="relative w-2/5 flex min-h-screen flex-col max-h-screen justify-center overflow-hidden py-6">
      {#if regex.test(fen)}
      <div class="p-6 mb-6 rounded-xl shadow-xl bg-slate-300 overflow-auto">
      {#await promise}
        <p>...loading</p>
      {:then} 
        
        <img class="w-full rounded-md" src="/api?Fen={fen}" alt="does this even work?" />
      {/await}
    </div>
      {/if}
      <div class="relative rounded-2xl bg-zinc-100 px-8 pt-8 pb-8 shadow-xl ring-1 ring-gray-900/5">
        <div class="mx-auto flex flex-col justify-center">
          <label for="website-admin" class="mb-2 hidden text-sm font-medium text-gray-900"></label>
          <div class="flex rounded-full border border-gray-800 bg-gray-50">
            <span class="inline-flex items-center border-gray-800 rounded-r-lg rounded-l-full border-r-2 bg-gray-200 px-3 text-sm text-gray-900">
                {#if regex.test(fen)}
                Valid Fen!!
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1" stroke="green" class="pl-1 w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  {:else}
                  Fen no good
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1" stroke="red" class="w-6 pl-1 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  {/if}
            </span>
            <input bind:value="{fen}" type="text" class="block w-full min-w-0 flex-1 focus:border-none grow p-2.5 text-sm text-gray-900 " placeholder="Add Fen String" />
            <div class=" h-full bg-gray-200 rounded-r-full border-l-2 border-gray-800">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 m-2 h-6 bg-gray-200">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 01-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 011.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 00-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 01-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 00-3.375-3.375h-1.5a1.125 1.125 0 01-1.125-1.125v-1.5a3.375 3.375 0 00-3.375-3.375H9.75" />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  