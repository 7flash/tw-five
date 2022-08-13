package main

/*
#cgo CFLAGS: -I/wallet-core/include
#cgo LDFLAGS: -L/wallet-core/build -L/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
#include <TrustWalletCore/TWHDWallet.h>
*/
import "C"

import "fmt"
import "os"
import "unsafe"

func main() {
	existingMnemonicStr, mnemonicNotExistErr := os.ReadFile(".mnemonic")

	emptyPassphrase := C.TWStringCreateWithUTF8Bytes(C.CString(""))

	var address unsafe.Pointer

	if mnemonicNotExistErr != nil {
		wallet := C.TWHDWalletCreate(128, emptyPassphrase)
		address = C.TWHDWalletGetAddressForCoin(wallet, C.enum_TWCoinType(C.TWCoinTypePolkadot))
		mnemonic := C.TWHDWalletMnemonic(wallet)

		os.WriteFile(".mnemonic", []byte(C.GoString(C.TWStringUTF8Bytes(mnemonic))), 0777)
	
		fmt.Printf("created new wallet.. \n")
	} else {	
		mnemonic := C.TWStringCreateWithUTF8Bytes(C.CString(string(existingMnemonicStr)))
		wallet := C.TWHDWalletCreateWithMnemonic(mnemonic, emptyPassphrase)
		address = C.TWHDWalletGetAddressForCoin(wallet, C.enum_TWCoinType(C.TWCoinTypePolkadot))

		fmt.Printf("loaded existing wallet.. \n")	
	}

	fmt.Printf("wallet address: %s \n", C.GoString(C.TWStringUTF8Bytes(address)))
}