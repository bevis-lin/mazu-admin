import SentimenMintRequest from 0x2ebc7543c6a3f855 //testnet

pub fun main() : {UInt64: SentimenMintRequest.MintRequest} {
  return SentimenMintRequest.getAllRequests()
}