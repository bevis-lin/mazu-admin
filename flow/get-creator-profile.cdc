import SentimenCreator from 0x2ebc7543c6a3f855

pub fun main(address: Address): SentimenCreator.Creator? {
  return SentimenCreator.getCreatorProfleByAddress(address: address)
}