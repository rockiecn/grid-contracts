package common

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	alterStr            = "alter"
	permStr             = "perm"
	addTStr             = "addT"
	setGSStr            = "setGS"
	setGDescStr         = "setGDesc"
	setRSStr            = "setRS"
	setkpPStr           = "setP"    // set k/p pledge
	setRateStr          = "setRate" // set manageRate/taxRate
	setLevelStr         = "setLevel"
	setPledgePerBStr    = "setPledgePerB"
	setPeStr            = "setPeriod"      // set period
	setWdTLStr          = "setWPTimeLimit" // set withdraw-pledge-time-limit
	setQuitRTLStr       = "setQRTimeLimit" // set quit-role-time-limit
	banTStr             = "banT"
	activateStr         = "activate"
	setStr              = "set"
	addOwnerStr         = "add"
	setCtlPermissionStr = "setRole"
	// middleware
	AlterTrafficPriceStr    = "alterTrafficPrice"
	AlterSpacePriceStr      = "alterSpacePrice"
	AlterSpaceSellerStr     = "alterSpaceSeller"
	AlterTrafficSellerStr   = "alterTrafficSeller"
	AlterTrafficReceiverStr = "alterTrafficReceiver"
	AlterSpaceReceiverStr   = "alterSpaceReceiver"
)

func GetAlterHash(instanAddr, authAddr, addr common.Address, t uint8, nonce *big.Int) []byte {
	hash := crypto.Keccak256(instanAddr.Bytes(), []byte(alterStr), addr.Bytes(), []byte{t})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetAddTHash(ctl1Addr, authAddr, ercAddr common.Address, ismint uint8, nonce *big.Int) []byte {
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(addTStr), ercAddr.Bytes(), []byte{ismint})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetSetGSHash(ctl1Addr, authAddr common.Address, gi uint64, state uint8, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, gi)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(setGSStr), tmp8, []byte{state})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetSetGDescHash(ctl1Addr, authAddr common.Address, gi uint64, desc []byte, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, gi)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(setGDescStr), tmp8, desc)
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// set k/p pledge amount
func GetSetkpPHash(ctl1Addr, authAddr common.Address, gi uint64, kp, pp, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, gi)
	m := common.LeftPadBytes(kp.Bytes(), 32)
	n := common.LeftPadBytes(pp.Bytes(), 32)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(setkpPStr), tmp8, m, n)
	m = common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetSetRateHash(ctl1Addr, authAddr common.Address, gi uint64, mr, tr uint8, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, gi)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(setRateStr), tmp8, []byte{mr}, []byte{tr})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetSetLevelHash(ctl3Addr, authAddr common.Address, gi uint64, level uint8, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, gi)
	hash := crypto.Keccak256(ctl3Addr.Bytes(), []byte(setLevelStr), tmp8, []byte{level})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetSetPlePerBHash(ctl3Addr, authAddr common.Address, gi uint64, kpPerB, ppPerB, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, gi)
	m := common.LeftPadBytes(kpPerB.Bytes(), 32)
	n := common.LeftPadBytes(ppPerB.Bytes(), 32)
	hash := crypto.Keccak256(ctl3Addr.Bytes(), []byte(setPledgePerBStr), tmp8, m, n)
	m = common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetSetPeriodHash(ctl3Addr, authAddr common.Address, pe uint64, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, pe)
	hash := crypto.Keccak256(ctl3Addr.Bytes(), []byte(setPeStr), tmp8)
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// get the hash of setting unpledgeTimeLimit
func GetSetUnpleTLHash(ctl3Addr, authAddr common.Address, timeLimit, nonce *big.Int) []byte {
	m := common.LeftPadBytes(timeLimit.Bytes(), 32)
	hash := crypto.Keccak256(ctl3Addr.Bytes(), []byte(setWdTLStr), m)
	m = common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// get the hash of setting quitRoleTimeLimit
func GetSetQuitRoleTLHash(ctl3Addr, authAddr common.Address, timeLimit, nonce *big.Int) []byte {
	m := common.LeftPadBytes(timeLimit.Bytes(), 32)
	hash := crypto.Keccak256(ctl3Addr.Bytes(), []byte(setQuitRTLStr), m)
	m = common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// activate keeper
func GetActivateHash(ctl1Addr, authAddr common.Address, ki uint64, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, ki)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(activateStr), tmp8)
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// abi.encodePacked(address(this), nonce, "set", account, isSet)
// set access
func GetSetHash(access, acc common.Address, nonce *big.Int, isSet bool) []byte {
	set := uint8(0)
	if isSet {
		set = 1
	}
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash := crypto.Keccak256(access.Bytes(), m, []byte(setStr), acc.Bytes(), []byte{set})
	return hash
}

// abi.encodePacked(uIndex,pIndex,nonce,start,end,size,tIndex,sprice)
func GetAddOrderHash(uIndex, pIndex, start, end, size, nonce uint64, tIndex uint8, sprice *big.Int) []byte {
	tmp1 := make([]byte, 8)
	tmp2 := make([]byte, 8)
	tmp3 := make([]byte, 8)
	tmp4 := make([]byte, 8)
	tmp5 := make([]byte, 8)
	tmp6 := make([]byte, 8)

	binary.BigEndian.PutUint64(tmp1, uIndex)
	binary.BigEndian.PutUint64(tmp2, pIndex)
	binary.BigEndian.PutUint64(tmp3, nonce)
	binary.BigEndian.PutUint64(tmp4, start)
	binary.BigEndian.PutUint64(tmp5, end)
	binary.BigEndian.PutUint64(tmp6, size)

	m := common.LeftPadBytes(sprice.Bytes(), 32)

	hash := crypto.Keccak256(tmp1, tmp2, tmp3, tmp4, tmp5, tmp6, []byte{tIndex}, m)

	return hash
}

// abi.encodePacked(address(this), "add", _a, isSet)
// abi.encodePacked(address(this), nonce, "perm", ha)
func GetAddOwnerHash(addr, ownerContractAddr, authAddr common.Address, isSet bool, nonce *big.Int) []byte {
	set := uint8(0)
	if isSet {
		set = 1
	}
	hash := crypto.Keccak256(addr.Bytes(), []byte(addOwnerStr), ownerContractAddr.Bytes(), []byte{set})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetSetCtlPermissionHash(contractAddr, grantedAcc, authAddr common.Address, role uint8, nonce *big.Int) []byte {
	hash := crypto.Keccak256(contractAddr.Bytes(), []byte(setCtlPermissionStr), grantedAcc.Bytes(), []byte{role})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// abi.encodePacked(ps.pIndex,ps.tIndex,ps.pay,ps.lost)
func GetProWithdrawHash(pindex uint64, tindex uint8, pay, lost *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, pindex)

	m := common.LeftPadBytes(pay.Bytes(), 32)
	n := common.LeftPadBytes(lost.Bytes(), 32)

	hash := crypto.Keccak256(tmp8, []byte{tindex}, m, n)
	return hash
}

func GetSetRSHash(ctl1Addr, authAddr common.Address, ri uint64, state uint8, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, ri)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(setRSStr), tmp8, []byte{state})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// GetBanTHash tokenAddr: erc20-token address
func GetBanTHash(ctl1Addr, authAddr, tokenAddr common.Address, isBan uint8, nonce *big.Int) []byte {
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(banTStr), tokenAddr.Bytes(), []byte{isBan})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// ----------------------------------middleware---------------------------------

// get the message hash for sending 'cashTrafficCheck' and 'cashSpaceCheck' transaction
// hash = keccak256(abi.encodePacked(address(ReadPay), nonce, sizeByte, seller))
func GetCashCheckHash(payAddr, sellerAddr common.Address, sizeByte uint64, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, sizeByte)

	m := common.LeftPadBytes(nonce.Bytes(), 32)

	hash := crypto.Keccak256(payAddr.Bytes(), m, tmp8, sellerAddr.Bytes())
	return hash
}

// used to get hash of calling 'alterTrafficPrice' and 'alterSpacePrice'
// param 'mwCtlAddr' is equal to middleware-control contract address
// param 'method' is equal to alterTrafficPriceStr or alterSpacePriceStr
// hash = keccak256(abi.encodePacked(address(middleware-control), "alterTrafficPrice", price))
// hash = keccak256(abi.encodePacked(address(middleware-control), "alterTrafficPrice", price))
func GetAlterPriceHash(mwCtlAddr, authAddr common.Address, newPrice uint64, nonce *big.Int, method string) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, newPrice)

	hash := crypto.Keccak256(mwCtlAddr.Bytes(), []byte(method), tmp8)

	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// used to get hash of calling 'alterTrafficSeller' 'alterSpaceSeller' 'alterSpaceReceiver' and 'alterTrafficReceiver'
// param 'mwCtlAddr' is equal to middleware-control contract address
// param 'method' is equal to alterTrafficSellerStr or alterSpaceSellerStr or alterTrafficReceiverStr or alterSpaceReceiverStr
func GetAlterSellerReceiverHash(mwCtlAddr, authAddr, newAddr common.Address, nonce *big.Int, method string) []byte {
	hash := crypto.Keccak256(mwCtlAddr.Bytes(), []byte(method), newAddr.Bytes())

	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func Sign(hash []byte, sk string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(sk)
	if err != nil {
		return nil, err
	}

	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func GetSigns(hash []byte, sks [5]string) [5][]byte {
	res := [5][]byte{}

	for i := 0; i < 5; i++ {
		sign, err := Sign(hash, sks[i])
		if err != nil {
			fmt.Println("get signs i: ", i, " err:", err)
			res[i] = nil
		} else {
			res[i] = sign
		}
	}

	return res
}
