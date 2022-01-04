import Sentimen from 0x78e84183b7e33d61
import SentimenAdmin from 0x78e84183b7e33d61
import SentimenMintRequest from 0x78e84183b7e33d61


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