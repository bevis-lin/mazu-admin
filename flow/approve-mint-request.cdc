import Sentimen from 0x2ebc7543c6a3f855
import SentimenAdmin from 0x2ebc7543c6a3f855
import SentimenMintRequest from 0x2ebc7543c6a3f855


transaction(requestId: UInt64) {

  let adminRef: &SentimenAdmin.Admin
  let minter: &Sentimen.NFTMinter

  prepare(signer:AuthAccount){
      self.adminRef = signer.borrow<&SentimenAdmin.Admin>(from: /storage/sentimenAdmin)?? panic("Could not borrow a reference to the SentimenAdmin")
      self.minter = signer.borrow<&Sentimen.NFTMinter>(from: /storage/NFTMinter)
            ?? panic("Could not borrow a reference to the NFT minter")
  }

  execute{
    SentimenMintRequest.approveMintRequest(adminRef: self.adminRef, minter: self.minter, mintRequestId: requestId)
  }

}