// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Eludius18BlockchainEludius18Blockchain from './eludius18blockchain.eludius18blockchain'


export default { 
  Eludius18BlockchainEludius18Blockchain: load(Eludius18BlockchainEludius18Blockchain, 'eludius18blockchain.eludius18blockchain'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}