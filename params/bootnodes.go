// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Akroma network.
var MainnetBootnodes = []string{
	"enode://2b2895b069aac74f1289ddbf6c083df7f549817fb488065261d9b5b8c4c175bc9103670ed32b70f3f7110a3658dfe9fbe39ff8828d9fc48df0a777971a4fc2b3@51.38.98.51:30303",
	"enode://bfe245cebd3f028880ea6844687baa8915475826ab74d9daf14bdbb6e931d35a30ebf6f78a5c66aff3ac34e3501e5d6ddf1e182ee3ce65907e84bac921133013@51.38.133.221:30303",
	"enode://d399460ab28039d59b7d791aea370131a03e2044f4164afb421173e4ebef37f926392574183901803ce6339dc535775b353daa13bc50fbd3bed073094ab790fb@54.38.215.141:30303",
	"enode://9249d309e36e7332fe1dc3acc557a72b098740aeab334322968c2cbffd0a86e12c1cb9bd11f027b1720a227eaffc35a7e0f955e5ccc730375cd2599909776b73@51.75.122.229:30303",
	"enode://492e7b3c922603102b75d090294b44ffb69d99b48c9aa96bfcda0f9e53092a619929564677cb771577300d6ab53723b18332db027a233eef64cdc761810c4c9f@207.246.91.9:30303",
	"enode://2254408a3093b28cbf56176de97a658cbfca05acdb607a24b7ae9341447e74fe1f295ff4549e4c26d28db5c69fb71e9a91b12827285318e68a3077a67525cb7b@45.32.228.212:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://7fdacaac41d84d56c0713797f75392c74419c56478e39e9831ecc3685e752dd53373df515b5caf1cb9b62af6682c13c2417be8dc059e3dae8ade9660409bd3f3@159.89.41.132:30303", //west-us
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
