package Security

import(
	"fmt"
	"net/http"
	"log"
	// "io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func GetPrimaryTrustChain(w http.ResponseWriter, r *http.Request){
	header:=GetPrimaryTrustChainHeader{}
	util.GetHeader(r,&header)

	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)

	fmt.Print("GetPrimaryTrustChainHeader:")
	fmt.Println(header)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	primaryTrustChainList,chainName,chainId,updatedAt:=GetPrimaryTrustChainDB(userId)

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&GetPrimaryTrustChainResponse{Code:200,Message:msg,ChainName:chainName,ChainId:chainId,UpdatedAt:updatedAt,PrimaryTrustChainList:primaryTrustChainList}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}