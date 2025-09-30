# DomRobot XML-RPC, JSON-RPC API Documentation

> This documentation describes the communication between client and API interface.

---

## Table of Contents

- [Chapter 1. Overview](#overview)
  - [1.1. XML format for API requests](#id2)
- [1.2. JSON format for API requests](#id484)
- [1.3. XML response format](#id512)
- [1.4. JSON response format](#id547)
- [1.5. Client implementations](#id582)
- [Chapter 2. Methods](#methods)
  - [2.1. Account](#id42)
      - [2.1.1. account.addrole](#account.addrole)
      - [2.1.2. account.changepassword](#account.changepassword)
      - [2.1.3. account.check](#account.check)
      - [2.1.4. account.create](#account.create)
      - [2.1.5. account.delete](#account.delete)
      - [2.1.6. account.getroles](#account.getroles)
      - [2.1.7. account.info](#account.info)
      - [2.1.8. account.list](#account.list)
      - [2.1.9. account.login](#account.login)
      - [2.1.10. account.logout](#account.logout)
      - [2.1.11. account.removerole](#account.removerole)
      - [2.1.12. account.unlock](#account.unlock)
      - [2.1.13. account.update](#account.update)
- [2.2. Accounting](#id2066)
    - [2.2.1. accounting.accountBalance](#accounting.accountBalance)
    - [2.2.2. accounting.creditLogById](#accounting.creditLogById)
    - [2.2.3. accounting.getInvoice](#accounting.getInvoice)
    - [2.2.4. accounting.getPreviewInvoice](#accounting.getPreviewInvoice)
    - [2.2.5. accounting.getReceipt](#accounting.getReceipt)
    - [2.2.6. accounting.getstatement](#accounting.getstatement)
    - [2.2.7. accounting.listInvoices](#accounting.listInvoices)
    - [2.2.8. accounting.lockedFunds](#accounting.lockedFunds)
    - [2.2.9. accounting.log](#accounting.log)
    - [2.2.10. accounting.refund](#accounting.refund)
    - [2.2.11. accounting.sendbypost](#accounting.sendbypost)
- [2.3. Application](#id2904)
    - [2.3.1. application.check](#application.check)
    - [2.3.2. application.create](#application.create)
    - [2.3.3. application.delete](#application.delete)
    - [2.3.4. application.info](#application.info)
    - [2.3.5. application.list](#application.list)
    - [2.3.6. application.update](#application.update)
- [2.4. Authinfo2](#id3639)
    - [2.4.1. authinfo2.create](#authinfo2.create)
    - [2.4.2. authinfo2.getprice](#authinfo2.getprice)
- [2.5. Certificate](#id3766)
    - [2.5.1. certificate.cancel](#certificate.cancel)
    - [2.5.2. certificate.create](#certificate.create)
    - [2.5.3. certificate.getProduct](#certificate.getProduct)
    - [2.5.4. certificate.info](#certificate.info)
    - [2.5.5. certificate.list](#certificate.list)
    - [2.5.6. certificate.listProducts](#certificate.listProducts)
    - [2.5.7. certificate.listRemainingNeededData](#certificate.listRemainingNeededData)
    - [2.5.8. certificate.log](#certificate.log)
    - [2.5.9. certificate.reissue](#certificate.reissue)
    - [2.5.10. certificate.renew](#certificate.renew)
    - [2.5.11. certificate.setAutorenew](#certificate.setAutorenew)
    - [2.5.12. certificate.setresendapproval](#certificate.setresendapproval)
    - [2.5.13. certificate.updateOrder](#certificate.updateOrder)
- [2.6. Contact](#id5242)
    - [2.6.1. contact.create](#contact.create)
    - [2.6.2. contact.delete](#contact.delete)
    - [2.6.3. contact.info](#contact.info)
    - [2.6.4. contact.list](#contact.list)
    - [2.6.5. contact.log](#contact.log)
    - [2.6.6. contact.sendbulkverification](#contact.sendbulkverification)
    - [2.6.7. contact.sendcontactverification](#contact.sendcontactverification)
    - [2.6.8. contact.update](#contact.update)
- [2.7. Customer](#id6062)
    - [2.7.1. customer.contactverificationsettingsinfo](#customer.contactverificationsettingsinfo)
    - [2.7.2. customer.contactverificationsettingsupdate](#customer.contactverificationsettingsupdate)
    - [2.7.3. customer.delete](#customer.delete)
    - [2.7.4. customer.info](#customer.info)
    - [2.7.5. customer.listdownloads](#customer.listdownloads)
    - [2.7.6. customer.requestdataexport](#customer.requestdataexport)
    - [2.7.7. customer.update](#customer.update)
- [2.8. Dnssec](#id7052)
    - [2.8.1. dnssec.adddnskey](#dnssec.adddnskey)
    - [2.8.2. dnssec.deleteall](#dnssec.deleteall)
    - [2.8.3. dnssec.deletednskey](#dnssec.deletednskey)
    - [2.8.4. dnssec.disablednssec](#dnssec.disablednssec)
    - [2.8.5. dnssec.enablednssec](#dnssec.enablednssec)
    - [2.8.6. dnssec.info](#dnssec.info)
    - [2.8.7. dnssec.listkeys](#dnssec.listkeys)
- [2.9. Domain](#id7530)
    - [2.9.1. domain.check](#domain.check)
    - [2.9.2. domain.create](#domain.create)
    - [2.9.3. domain.delete](#domain.delete)
    - [2.9.4. domain.getalldomainprices](#domain.getalldomainprices)
    - [2.9.5. domain.getdomainprice](#domain.getdomainprice)
    - [2.9.6. domain.getextradatarules](#domain.getextradatarules)
    - [2.9.7. domain.getPrices](#domain.getPrices)
    - [2.9.8. domain.getPromos](#domain.getPromos)
    - [2.9.9. domain.getRules](#domain.getRules)
    - [2.9.10. domain.getTldGroups](#domain.getTldGroups)
    - [2.9.11. domain.info](#domain.info)
    - [2.9.12. domain.list](#domain.list)
    - [2.9.13. domain.log](#domain.log)
    - [2.9.14. domain.priceChanges](#domain.priceChanges)
    - [2.9.15. domain.push](#domain.push)
    - [2.9.16. domain.removeClientHold](#domain.removeClientHold)
    - [2.9.17. domain.renew](#domain.renew)
    - [2.9.18. domain.restore](#domain.restore)
    - [2.9.19. domain.setClientHold](#domain.setClientHold)
    - [2.9.20. domain.stats](#domain.stats)
    - [2.9.21. domain.trade](#domain.trade)
    - [2.9.22. domain.transfer](#domain.transfer)
    - [2.9.23. domain.transfercancel](#domain.transfercancel)
    - [2.9.24. domain.transferOut](#domain.transferOut)
    - [2.9.25. domain.update](#domain.update)
    - [2.9.26. domain.whois](#domain.whois)
- [2.10. Dyndns](#id11241)
    - [2.10.1. dyndns.changepassword](#dyndns.changepassword)
    - [2.10.2. dyndns.check](#dyndns.check)
    - [2.10.3. dyndns.create](#dyndns.create)
    - [2.10.4. dyndns.delete](#dyndns.delete)
    - [2.10.5. dyndns.info](#dyndns.info)
    - [2.10.6. dyndns.list](#dyndns.list)
    - [2.10.7. dyndns.log](#dyndns.log)
    - [2.10.8. dyndns.updateRecord](#dyndns.updateRecord)
- [2.11. DyndnsSubscription](#id11728)
    - [2.11.1. dyndnssubscription.cancel](#dyndnssubscription.cancel)
    - [2.11.2. dyndnssubscription.create](#dyndnssubscription.create)
    - [2.11.3. dyndnssubscription.list](#dyndnssubscription.list)
    - [2.11.4. dyndnssubscription.listProducts](#dyndnssubscription.listProducts)
- [2.12. Host](#id11976)
    - [2.12.1. host.check](#host.check)
    - [2.12.2. host.create](#host.create)
    - [2.12.3. host.delete](#host.delete)
    - [2.12.4. host.info](#host.info)
    - [2.12.5. host.list](#host.list)
    - [2.12.6. host.update](#host.update)
- [2.13. Hosting](#id12347)
    - [2.13.1. hosting.cancel](#hosting.cancel)
    - [2.13.2. hosting.controlPanel](#hosting.controlPanel)
    - [2.13.3. hosting.create](#hosting.create)
    - [2.13.4. hosting.getPrices](#hosting.getPrices)
    - [2.13.5. hosting.issuspended](#hosting.issuspended)
    - [2.13.6. hosting.list](#hosting.list)
    - [2.13.7. hosting.reinstate](#hosting.reinstate)
    - [2.13.8. hosting.unsuspend](#hosting.unsuspend)
    - [2.13.9. hosting.updatePeriod](#hosting.updatePeriod)
- [2.14. Message](#id12887)
    - [2.14.1. message.ack](#message.ack)
    - [2.14.2. message.poll](#message.poll)
- [2.15. Nameserver](#id12987)
    - [2.15.1. nameserver.check](#nameserver.check)
    - [2.15.2. nameserver.clone](#nameserver.clone)
    - [2.15.3. nameserver.create](#nameserver.create)
    - [2.15.4. nameserver.createRecord](#nameserver.createRecord)
    - [2.15.5. nameserver.delete](#nameserver.delete)
    - [2.15.6. nameserver.deleteRecord](#nameserver.deleteRecord)
    - [2.15.7. nameserver.export](#nameserver.export)
    - [2.15.8. nameserver.exportlist](#nameserver.exportlist)
    - [2.15.9. nameserver.exportrecords](#nameserver.exportrecords)
    - [2.15.10. nameserver.info](#nameserver.info)
    - [2.15.11. nameserver.list](#nameserver.list)
    - [2.15.12. nameserver.update](#nameserver.update)
    - [2.15.13. nameserver.updateRecord](#nameserver.updateRecord)
- [2.16. NameserverSet](#id14267)
    - [2.16.1. nameserverset.create](#nameserverset.create)
    - [2.16.2. nameserverset.delete](#nameserverset.delete)
    - [2.16.3. nameserverset.info](#nameserverset.info)
    - [2.16.4. nameserverset.list](#nameserverset.list)
    - [2.16.5. nameserverset.update](#nameserverset.update)
- [2.17. News](#id14777)
    - [2.17.1. news.list](#news.list)
- [2.18. Nichandle](#id14917)
    - [2.18.1. nichandle.list](#nichandle.list)
- [2.19. Pdf](#id15018)
    - [2.19.1. pdf.document](#pdf.document)
    - [2.19.2. pdf.get](#pdf.get)
- [2.20. Tag](#id15125)
    - [2.20.1. tag.create](#tag.create)
    - [2.20.2. tag.delete](#tag.delete)
    - [2.20.3. tag.info](#tag.info)
    - [2.20.4. tag.list](#tag.list)
    - [2.20.5. tag.update](#tag.update)
- [Chapter 3. Datatypes](#types)
    - [3.1. _true](#type._true)
    - [3.2. addressTitle](#type.addresstitle)
    - [3.3. applicationOrder](#type.applicationorder)
    - [3.4. array](#type.array)
    - [3.5. array_domain](#type.array_domain)
    - [3.6. array_float](#type.array_float)
    - [3.7. array_identifier](#type.array_identifier)
    - [3.8. array_int](#type.array_int)
    - [3.9. array_ip](#type.array_ip)
    - [3.10. array_text](#type.array_text)
    - [3.11. array_text255](#type.array_text255)
    - [3.12. array_text64](#type.array_text64)
    - [3.13. auDomainIdType](#type.audomainidtype)
    - [3.14. auDomainRelation](#type.audomainrelation)
    - [3.15. auDomainRelationType](#type.audomainrelationtype)
    - [3.16. auEligibilityIdType](#type.aueligibilityidtype)
    - [3.17. base64](#type.base64)
    - [3.18. bgAppNumber](#type.bgappnumber)
    - [3.19. boolean](#type.boolean)
    - [3.20. boolean_3](#type.boolean_3)
    - [3.21. brCnpj](#type.brcnpj)
    - [3.22. brCpf](#type.brcpf)
    - [3.23. caLegalType](#type.calegaltype)
    - [3.24. contact](#type.contact)
    - [3.25. contactOrder](#type.contactorder)
    - [3.26. contactType](#type.contacttype)
    - [3.27. country](#type.country)
    - [3.28. customercurrency](#type.customercurrency)
    - [3.29. date](#type.date)
    - [3.30. dateTime](#type.datetime)
    - [3.31. dnskey](#type.dnskey)
    - [3.32. dnssecAlgorithm](#type.dnssecalgorithm)
    - [3.33. dnssecDigestType](#type.dnssecdigesttype)
    - [3.34. dnssecDomainStatus](#type.dnssecdomainstatus)
    - [3.35. dnssecFlag](#type.dnssecflag)
    - [3.36. dnssecKeyStatus](#type.dnsseckeystatus)
    - [3.37. documentformat](#type.documentformat)
    - [3.38. domainLogOrder](#type.domainlogorder)
    - [3.39. domainOrder](#type.domainorder)
    - [3.40. ds](#type.ds)
    - [3.41. dunsNumber](#type.dunsnumber)
    - [3.42. eea_countries](#type.eea_countries)
    - [3.43. email](#type.email)
    - [3.44. emailoptional](#type.emailoptional)
    - [3.45. esIdType](#type.esidtype)
    - [3.46. esLegalForm](#type.eslegalform)
    - [3.47. es_nif_nie](#type.es_nif_nie)
    - [3.48. extData](#type.extdata)
    - [3.49. fiHenkilotunnus](#type.fihenkilotunnus)
    - [3.50. float](#type.float)
    - [3.51. float_signed](#type.float_signed)
    - [3.52. hkIndustryType](#type.hkindustrytype)
    - [3.53. hostname](#type.hostname)
    - [3.54. hrOib](#type.hroib)
    - [3.55. ieHolderType](#type.ieholdertype)
    - [3.56. int](#type.int)
    - [3.57. ip](#type.ip)
    - [3.58. ipList](#type.iplist)
    - [3.59. ip_url](#type.ip_url)
    - [3.60. irCompanyRegistrationType](#type.ircompanyregistrationtype)
    - [3.61. irNationalId](#type.irnationalid)
    - [3.62. irOrganizationId](#type.irorganizationid)
    - [3.63. isKennitala](#type.iskennitala)
    - [3.64. itCodiceFiscale](#type.itcodicefiscale)
    - [3.65. krCtfyType](#type.krctfytype)
    - [3.66. language](#type.language)
    - [3.67. message_status](#type.message_status)
    - [3.68. message_type](#type.message_type)
    - [3.69. myOrgType](#type.myorgtype)
    - [3.70. noPersonIdentifier](#type.nopersonidentifier)
    - [3.71. nsList](#type.nslist)
    - [3.72. nsSetType](#type.nssettype)
    - [3.73. nsType](#type.nstype)
    - [3.74. password](#type.password)
    - [3.75. paymentType](#type.paymenttype)
    - [3.76. period](#type.period)
    - [3.77. phone](#type.phone)
    - [3.78. phoneoptional](#type.phoneoptional)
    - [3.79. ptLegitimacy](#type.ptlegitimacy)
    - [3.80. ptRegistrationBasis](#type.ptregistrationbasis)
    - [3.81. recordType](#type.recordtype)
    - [3.82. remarks](#type.remarks)
    - [3.83. renewalMode](#type.renewalmode)
    - [3.84. seIdNo](#type.seidno)
    - [3.85. signMethod](#type.signmethod)
    - [3.86. skLegalForm](#type.sklegalform)
    - [3.87. swissuid](#type.swissuid)
    - [3.88. swissupi](#type.swissupi)
    - [3.89. tagUpdateAdd](#type.tagupdateadd)
    - [3.90. tagUpdateRem](#type.tagupdaterem)
    - [3.91. text](#type.text)
    - [3.92. text0](#type.text0)
    - [3.93. text0100](#type.text0100)
    - [3.94. text0255](#type.text0255)
    - [3.95. text064](#type.text064)
    - [3.96. text10](#type.text10)
    - [3.97. text1024](#type.text1024)
    - [3.98. text255](#type.text255)
    - [3.99. text64](#type.text64)
    - [3.100. tfaMethod](#type.tfamethod)
    - [3.101. timestamp](#type.timestamp)
    - [3.102. token0255](#type.token0255)
    - [3.103. token255](#type.token255)
    - [3.104. trCitizenId](#type.trcitizenid)
    - [3.105. transferAnswer](#type.transferanswer)
    - [3.106. transferMode](#type.transfermode)
    - [3.107. uaTrademarkType](#type.uatrademarktype)
    - [3.108. urlRedirectType](#type.urlredirecttype)
    - [3.109. usCategory](#type.uscategory)
    - [3.110. usPurpose](#type.uspurpose)
    - [3.111. username](#type.username)
    - [3.112. vatNo](#type.vatno)
    - [3.113. vatNoInternational](#type.vatnointernational)
    - [3.114. zuerichuid](#type.zuerichuid)
- [Chapter 4. Result Codes](#errorcodes)

---


## Chapter 1: Overview


## Chapter 1. Overview {#overview}

The communication is realized through XML-RPC and JSON-RPC (Remote Procedure Call) services. This is a specification by sending HTTP(S) requests to a assigned address. All client requests and server answers are using the XML or JSON format. The client must send the XML or JSON request via HTTP(S) POST and the server response is a XML or JSON document as well.

The first request must be an account.login command. In case of a succeded login you will receive a cookie session id in the header of the response. Please send it for all further requests.


### 1.1. XML format for API requests {#id2}

The addresses of the provided services are:

- for the test environment https://api.ote.domrobot.com/xmlrpc/
- for the production environment https://api.domrobot.com/xmlrpc/

For more information about XML-RPC, visit the web site at http://www.xmlrpc.com/.

Example of a XML-RPC request:


```xml
<?xml version="1.0" encoding="UTF-8"?>
        <methodCall>
          <methodName>account.login</methodName>
            <params>
              <param>
                <value>
                  <struct>
                    <member>
                      <name>user</name>
                        <value>
                          <string>your_username</string>
                        </value>
                    </member>
                    <member>
                      <name>pass</name>
                        <value>
                          <string>your_password</string>
                        </value>
                    </member>
                    <member>
                     <name>lang</name>
                      <value>
                        <string>en</string>
                      </value>
                    </member>
                    <member>
                      <name>clTRID</name>
                        <value>
                          <string>CLIENT-123123</string>
                        </value>
                    </member>
                  </struct>
                </value>
              </param>
             </params>
        </methodCall>
          ```

Main (optional) parameters

- lang: Language of the return message ('en' or 'de')
- clTRID: The clTRID stands for Client Transaction Identifier and may be helpful for your support team

Method parameters

Are described in Chapter 2: Methods.


## 1.2. JSON format for API requests {#id484}

The addresses of the provided services are:

- for the test environment https://api.ote.domrobot.com/jsonrpc/
- for the production environment https://api.domrobot.com/jsonrpc/

For more information about JSON-RPC, visit the web site at https://www.jsonrpc.org/.

Example of a JSON-RPC request:


```json

      {
  	"method": "account.login",
  	"params": {"user":"blakeks", "pass":"SCHEDULED", "lang":"de", "clTRID":"CLIENT-123123"}
      }
      ```

Main (optional) parameters

- lang: Language of the return message ('en' or 'de')
- clTRID: The clTRID stands for Client Transaction Identifier and may be helpful for your support team

Method parameters

Are described in Chapter 2: Methods.


## 1.3. XML response format {#id512}

Example of a XML-RPC success response:


```xml
<?xml version="1.0" encoding="UTF-8"?>
        <methodResponse>
            <params>
                <param>
                    <value>
                        <struct>
                            <member>
                                <name>code</name>
                                <value>
                                    <int>1000</int>
                                </value>
                            </member>
                            <member>
                                <name>msg</name>
                                <value>
                                    <string>Command completed successfully</string>
                                </value>
                            </member>
                            <member>
                                <name>resData</name>
                                <value>
                                    <struct>
                                        <member>
                                            <name>customerId</name>
                                            <value>
                                                <int>10069</int>
                                            </value>
                                        </member>
                                        <member>
                                            <name>customerNo</name>
                                            <value>
                                                <int>110069</int>
                                            </value>
                                        </member>
                                        <member>
                                            <name>accountId</name>
                                            <value>
                                                <int>61140</int>
                                            </value>
                                        </member>
                                        <member>
                                            <name>tfa</name>
                                            <value>
                                                <string>0</string>
                                            </value>
                                        </member>
                                    </struct>
                                </value>
                            </member>
                            <member>
                                <name>svTRID</name>
                                <value>
                                    <string>20230210-38566814-ote</string>
                                </value>
                            </member>
                            <member>
                                <name>runtime</name>
                                <value>
                                    <double>0.058500</double>
                                </value>
                            </member>
                        </struct>
                    </value>
                </param>
            </params>
        </methodResponse>
        ```

Example of a XML-RPC error response:


```xml
<?xml version="1.0" encoding="UTF-8"?>
        <methodResponse>
            <params>
                <param>
                    <value>
                        <struct>
                            <member>
                                <name>code</name>
                                <value>
                                    <int>2003</int>
                                </value>
                            </member>
                            <member>
                                <name>msg</name>
                                <value>
                                    <string>Required parameter missing</string>
                                </value>
                            </member>
                            <member>
                                <name>reasonCode</name>
                                <value>
                                    <string>MISSING_ID</string>
                                </value>
                            </member>
                            <member>
                                <name>reason</name>
                                <value>
                                    <string>The following parameter is missing: id</string>
                                </value>
                            </member>
                            <member>
                                <name>details</name>
                                <value>
                                    <array>
                                        <data>
                                            <value>
                                                <struct>
                                                    <member>
                                                        <name>code</name>
                                                        <value>
                                                            <string>PARAM_MISSING</string>
                                                        </value>
                                                    </member>
                                                    <member>
                                                        <name>msg</name>
                                                        <value>
                                                            <string>The parameter 'id' is missing</string>
                                                        </value>
                                                    </member>
                                                </struct>
                                            </value>
                                        </data>
                                    </array>
                                </value>
                            </member>
                            <member>
                                <name>svTRID</name>
                                <value>
                                    <string>20230210-38566799-ote</string>
                                </value>
                            </member>
                            <member>
                                <name>runtime</name>
                                <value>
                                    <double>0.004500</double>
                                </value>
                            </member>
                        </struct>
                    </value>
                </param>
            </params>
        </methodResponse>
        ```

Return parameters:

- code: Return code (described in Chapter 4: Result Codes)
- msg: Return message (described in Chapter 4: Result Codes)
- reasonCode: Additional short error message tag
- reason: Additional error message
- resData: Data result values
- svTRID: The svTRID stands for Server Transaction Identifier and may be helpful if you contact our support team


## 1.4. JSON response format {#id547}

Example of a JSON-RPC success response:


```json

          {
            "code": 1000,
            "msg": "Command completed successfully",
            "resData": {
              "customerId": 10069,
              "customerNo": 110069,
              "accountId": 61140,
              "tfa": "0"
            }
          }
    ```

Example of a JSON-RPC error response:


```json

          {
            "code": 2003,
            "msg": "Required parameter missing",
            "reasonCode": "MISSING_ID",
            "reason": "The following parameter is missing: id",
            "details": [{ "code": "PARAM_MISSING", "msg": "The parameter 'id' is missing" }]
          }
    ```

Return parameters:

- code: Return code (described in Chapter 4: Result Codes)
- msg: Return message (described in Chapter 4: Result Codes)
- reasonCode: Additional short error message tag
- reason: Additional error message
- resData: Data result values
- svTRID: The svTRID stands for Server Transaction Identifier and may be helpful if you contact our support team


## 1.5. Client implementations {#id582}

There are client implementations in different programming languages that are developed and supported by us. Examples on how to use these can be found in the README.md file of each project. There are even completely ready to use plugins for WHMCS and some other systems. You can find all these here: https://www.inwx.com/en/offer/api


## Chapter 2: Methods


## Chapter 2. Methods {#methods}


### 2.1. Account {#id42}

The account object provides methods concerning to your account.


##### 2.1.1. account.addrole {#account.addrole}

Add a role to an account


###### 2.1.1.1. Input {#id3}

Table 2.1. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| accountId | Id of the account | int | true |  |
| roleId | Id of the role (20000: Standard, 20001: Accounting, 20002: Domain, 20003: Hosting, 20004: DNS, 20005: Authcodes) | int | true |  |


###### 2.1.1.2. Output {#id4}

No additional return parameters


##### 2.1.2. account.changepassword {#account.changepassword}

Change current password.


###### 2.1.2.1. Input {#id6}

Table 2.2. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| username | Customers username | text64 | true |  |
| currentpassword | Throws EPP1200 if provided password doesn't matches current password | password | true |  |
| password | Customers password | password | true |  |
| testing | Execute command in testing mode | boolean | false | false |


###### 2.1.2.2. Output {#id7}

No additional return parameters


##### 2.1.3. account.check {#account.check}

Check whether you have currently a session.


###### 2.1.3.1. Input {#id9}

No parameters allowed


###### 2.1.3.2. Output {#id10}

Table 2.3. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| user | Your username | text64 |  |


##### 2.1.4. account.create {#account.create}

Create a new account.


###### 2.1.4.1. Input {#id12}

Table 2.4. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| username | accounts username | username | true |  |
| title | title of the account | addressTitle | true |  |
| firstname | accounts firstname | text64 | true |  |
| lastname | accounts lastname | text64 | true |  |
| street | accounts street | text255 | true |  |
| pc | accounts post code | text10 | true |  |
| city | accounts city | text64 | true |  |
| cc | accounts country code | country | true |  |
| email | accounts email address | email | true |  |
| org | accounts organisation/company | token0255 | false |  |
| voice | accounts phone number | phone | false |  |
| fax | accounts fax machine number | phoneoptional | false |  |
| www | accounts website | token0255 | false |  |
| language | interface language | language | false |  |


###### 2.1.4.2. Output {#id13}

Table 2.5. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the created account | int |  |


##### 2.1.5. account.delete {#account.delete}

Function for a main account to delete subaccounts


###### 2.1.5.1. Input {#id15}

Table 2.6. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| ids | One or more account id(s) to delete | array_int | true | false |


###### 2.1.5.2. Output {#id16}

Table 2.7. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| countDeleted | Count of successfully deleted accounts | int |  |


##### 2.1.6. account.getroles {#account.getroles}

Get roles


###### 2.1.6.1. Input {#id18}

Table 2.8. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| accountId | Id of the account | int | false |  |


###### 2.1.6.2. Output {#id19}

Table 2.9. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roles | List of the roles | array |  |


##### 2.1.7. account.info {#account.info}

Get your account details.


###### 2.1.7.1. Input {#id21}

Table 2.10. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| wide | More detailed output | int | false |  |


###### 2.1.7.2. Output {#id22}

Table 2.11. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| accountId | Account id | int |  |
| customerId | Customer id | int |  |
| customerNo | Your unique Customer number | int |  |
| username | Username | text64 |  |
| title | Salutation title | addressTitle |  |
| firstname | Customers firstname | text64 |  |
| lastname | Customers lastname | text64 |  |
| org | Customers organisation/company | text064 |  |
| street | Customers street | text255 |  |
| pc | Customers postal code | text10 |  |
| city | Customers city | text64 |  |
| cc | Customers country | country |  |
| voice | Customers phone number | phone |  |
| fax | Customers fax number | phoneoptional |  |
| www | Customers web address | token0255 |  |
| email | Customers email address | email |  |
| servicePin | Customers service pin | int |  |
| crDate | Customers date of account creation | dateTime |  |
| secureMode | Customers secure mode | boolean |  |
| signPdfs | Customer gets signed pdfs | boolean |  |
| summaryInvoice | Customer gets summary invoice | boolean |  |
| mailListId | List of subscribed mail-list-ids | array_int | Yes |
| language | Customers language | language |  |
| notificationEmail | Customer gets notification emails | boolean |  |
| notificationQueue | Customer uses notification queue | boolean |  |
| lowBalance | Customers low balance notification value | float |  |
| renewalReport | Customer gets renewal reports | boolean |  |
| paymentType | Customers type of payment | paymentType |  |
| bankAccHolder | Customers name of bank account holder | token255 |  |
| bankName | Customers name of bank | token255 |  |
| bankCode | Customers bank code | token255 |  |
| bankAccHolderNo | Customers bank account number | text64 |  |
| vat | Customers vat | int |  |
| vatNo | Customers company vat number | vatNo |  |
| whoisProvider | Default whois provider | token255 |  |
| whoisUrl | Default whois url | token255 |  |
| defaultRegistrant | Default registrant contact handle | int |  |
| defaultAdmin | Default administrative contact handle | int |  |
| defaultTech | Default technical contact handle | int |  |
| invoiceRouting | Invoice routing ID | text0100 |  |
| orderReference | Order Reference | text0100 |  |
| defaultBilling | Default billing contact handle | int |  |
| invoiceXml | Xml type invoices will be send | boolean |  |
| invoicePdf | Pdf type invoices will be send | boolean |  |
| defaultNsset | Default nameserver set | int |  |
| defaultWeb | Default web nameserver entry | token0255 | Yes |
| defaultMail | Default mail nameserver entry | text0255 | Yes |
| defaultImportNS | Import NS in case of Domain Transfer | boolean | Yes |
| defaultRenewalMode | Default domain renewal mode | renewalMode |  |
| lastLogin | Date of last login | dateTime |  |
| loginCount | Count of logins | int |  |
| rowsPerPage | Default rows per page value | int |  |
| tfa | 2-factor-authentification method | tfaMethod |  |
| lastIP | Customers ip address of last access | ip |  |
| verification | Customers data verfication pending flag | int |  |
| emailBilling | Customers email address for billing | email |  |
| emailAutomated | Email address for automated emails | email |  |
| currency | currency | customercurrency |  |
| isReseller | Defines if the customer is a reseller or not | boolean_3 |  |
| serviceProviderId | The provider-id of the customer | int |  |
| supplimentInvoiceText | Suppliment text to the invoice | text0100 |  |
| disablepremium | Are premium domains disabled by the customer? | boolean |  |
| enableAdvertising | Advertising is allowed on expired domains. | boolean |  |
| wdrpEmail | Email address for WDRP notifications | email |  |


##### 2.1.8. account.list {#account.list}

Function for a main account to get a list of all their subaccounts


###### 2.1.8.1. Input {#id24}

No parameters allowed


###### 2.1.8.2. Output {#id25}

Table 2.12. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| adminId | Admin account id | int |  |
| accounts | List of subaccounts | array |  |
| ... id | Account ID. | int |  |
| ... username | Account username. | text64 |  |
| ... firstname | Account first name. | text64 |  |
| ... lastname | Account last name. | text64 |  |
| ... org | Account organisation/company. | text064 |  |
| ... street | Account address, street. | text255 |  |
| ... pc | Account postal code. | text10 |  |
| ... city | Account city. | text64 |  |
| ... cc | Account country code. | country |  |
| ... voice | Account phone number. | phone |  |
| ... fax | Account fax number. | phoneoptional |  |
| ... email | Account email address. | email |  |
| ... tfaEnabled | 2-factor-authentification status. | boolean |  |


##### 2.1.9. account.login {#account.login}

Log in to API.


###### 2.1.9.1. Input {#id27}

Table 2.13. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| user | Your username | text64 | true |  |
| pass | Your password | password | true |  |
| case-insensitive | Case insensitivity for username | boolean | false | false |


###### 2.1.9.2. Output {#id28}

Table 2.14. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| customerId | Customers id | int |  |
| customerNo | Your unique Customer number | int |  |
| accountId | Customers account id | int |  |
| dyndns | DynDNS enabled flag | int |  |
| tfa | 2-factor-authentification method | tfaMethod |  |


##### 2.1.10. account.logout {#account.logout}

Log out from API.


###### 2.1.10.1. Input {#id30}

No parameters allowed


###### 2.1.10.2. Output {#id31}

No additional return parameters


##### 2.1.11. account.removerole {#account.removerole}

Remove role of an account


###### 2.1.11.1. Input {#id33}

Table 2.15. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| accountId | Id of the account | int | true |  |
| roleId | Id of the role | int | true |  |


###### 2.1.11.2. Output {#id34}

No additional return parameters


##### 2.1.12. account.unlock {#account.unlock}

Unlock your account


###### 2.1.12.1. Input {#id36}

Table 2.16. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| tan | Your TAN (Transaction number) | text10 | false |  |


###### 2.1.12.2. Output {#id37}

No additional return parameters


##### 2.1.13. account.update {#account.update}

Update customers account data.


###### 2.1.13.1. Input {#id39}

Table 2.17. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| username | Customers username | username | false |  |
| title | Customers salutation title | addressTitle | false |  |
| firstname | Customers firstname | text64 | false |  |
| lastname | Customers lastname | text64 | false |  |
| org | Customers organisation/company | token0255 | false |  |
| street | Customers street | text255 | false |  |
| pc | Customers postal code | text10 | false |  |
| city | Customers city | text64 | false |  |
| cc | Customers country | country | false |  |
| voice | Customers phone number | phone | false |  |
| fax | Customers fax number | phoneoptional | false |  |
| www | Customers web address | token0255 | false |  |
| email | Customers email address | email | false |  |
| emailAutomated | Customers automated email address | emailoptional | false |  |
| emailBilling | Customers email address for billing | emailoptional | false |  |
| renewalReportEmail | Customers renewal report email address | emailoptional | false |  |
| servicePin | Customers service pin | text64 | false |  |
| summaryInvoice | Customer gets summary invoice | boolean | false |  |
| mailListId | One or more mail-list-id(s) | array_int | false |  |
| language | Customers language | language | false |  |
| notificationEmail | Customer get notification emails | boolean | false |  |
| notificationQueue | Customer uses notification queue | boolean | false |  |
| renewalReport | Customer gets renewal reports emails | boolean | false |  |
| lowBalance | Customers low balance notification value | float_signed | false |  |
| bankAccHolder | Customers name of bank account holder | token255 | false |  |
| bankCode | Customers bank code | text64 | false |  |
| bankName | Customers name of bank | text64 | false |  |
| bankAccHolderNo | Customers bank account number | text64 | false |  |
| vatNo | Customers company vat number | vatNoInternational | false |  |
| whoisProvider | Default whois provider | token0255 | false |  |
| whoisUrl | Default whois url | token0255 | false |  |
| defaultRegistrant | Default registrant contact handle | int | false |  |
| defaultAdmin | Default admin contact handle | int | false |  |
| defaultTech | Default tech contact handle | int | false |  |
| invoiceRouting | Invoice routing ID | text0100 | false |  |
| orderReference | Order Reference | text0100 | false |  |
| defaultBilling | Default billing contact handle | int | false |  |
| invoiceXml | Xml type invoices will be send | boolean | false |  |
| invoicePdf | Pdf type invoices will be send | boolean | false |  |
| defaultNsset | Default nameserver set | int | false |  |
| defaultWeb | Default web nameserver entry | token0255 | false |  |
| defaultMail | Default mail nameserver entry | text0255 | false |  |
| defaultImportNS | Import NS in Case of Domain Transfer | boolean | false |  |
| defaultRenewalMode | Default domain renewal mode | renewalMode | false |  |
| supplimentinvoicetext | Suppliment text to the invoice | text0100 | false |  |
| password | Customers password | password | false |  |
| rowsPerPage | Default rows per page value | int | false |  |
| isReseller | Is Customer a reseller? | boolean_3 | false |  |
| testing | Execute command in testing mode | boolean | false | false |
| nif | NIF of a natural person or company (only for customers of InternetworX SLU) | text10 | false |  |
| idCard | Your ID card number (only for natural person with residence in the EU for InternetworX SLU) | text64 | false |  |
| disablepremium | Are premium domains disabled by the customer? | boolean | false |  |
| enableAdvertising | Advertising is allowed on expired domains. | boolean | false |  |
| wdrpEmail | Specify the email address for WDRP notifications | emailoptional | false |  |
| wdrpEmailName | Specify the sender name for WDRP notifications | text0255 | false |  |
| errpEmail | Specify the sender email address for ERRP notifications | emailoptional | false |  |
| errpEmailName | Specify the sender name for ERRP notifications | text0255 | false |  |
| foaEmail | Specify the sender email address for FOA notifications | emailoptional | false |  |
| foaEmailName | Specify the sender name for FOA notifications | text0255 | false |  |


###### 2.1.13.2. Output {#id40}

No additional return parameters


## 2.2. Accounting {#id2066}

The accounting object provides methods concerning to your account balance and invoices.


#### 2.2.1. accounting.accountBalance {#accounting.accountBalance}

Account balance details.


##### 2.2.1.1. Input {#id43}

No parameters allowed


##### 2.2.1.2. Output {#id44}

Table 2.18. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| total | Accumulated amount of received payments | float |  |
| available | Deposit that is available for transactions | float |  |
| locked | Deposit that is locked for transactions in process | float |  |
| creditLimit | Customers credit limit value | float |  |


#### 2.2.2. accounting.creditLogById {#accounting.creditLogById}

Get information about a specified credit log.


##### 2.2.2.1. Input {#id46}

Table 2.19. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| creditLogId | ID of the credit log to be shown | int | true |  |


##### 2.2.2.2. Output {#id47}

Table 2.20. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| log |  |  |  |
| ... id | ID of the credit log | int |  |
| ... amount | Amount value | float |  |
| ... creditTime | Log timestamp | dateTime |  |
| ... customerId | ID of the customer | int |  |
| ... creditType | Payment method of the deposit | text64 |  |
| ... transactionId | Transaction ID of the deposit | text64 |  |
| ... refundId | Associated ID of the refund (if it there is one) | int |  |
| ... transactionDetails | Details about the deposit (e.g. Paypal account) | text64 |  |
| ... currency | Currency of the deposit | text64 |  |
| ... last4 | The last four digits of the card | text64 |  |


#### 2.2.3. accounting.getInvoice {#accounting.getInvoice}

Returns invoice pdf document.


##### 2.2.3.1. Input {#id49}

Table 2.21. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| invoiceId | Id of the invoice | text64 | true |  |


##### 2.2.3.2. Output {#id50}

Table 2.22. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| pdf | Invoice pdf as base64 encoded | base64 |  |


#### 2.2.4. accounting.getPreviewInvoice {#accounting.getPreviewInvoice}

Returns preview invoice pdf document.


##### 2.2.4.1. Input {#id52}

Table 2.23. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| year | Year of the preview invoice | int | true |  |
| month | Month of the preview invoice | int | true |  |


##### 2.2.4.2. Output {#id53}

Table 2.24. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| pdf | Invoice pdf as base64 encoded | base64 |  |


#### 2.2.5. accounting.getReceipt {#accounting.getReceipt}

Returns receipt pdf document.


##### 2.2.5.1. Input {#id55}

Table 2.25. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | id of the payment | int | true |  |
| amount | amount to be receipt or refund | text64 | true |  |
| details | details required for pdf text | text64 | true |  |
| date | date of request | text64 | true |  |


##### 2.2.5.2. Output {#id56}

Table 2.26. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| pdf | Receipt/Refund PDF as base64 encoded | base64 |  |


#### 2.2.6. accounting.getstatement {#accounting.getstatement}

Statement PDF Document of customers transactions.


##### 2.2.6.1. Input {#id58}

Table 2.27. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| dateFrom | Log start date | timestamp | true |  |
| dateTo | Log end date | timestamp | true |  |
| format | Format of the requested document | documentformat | false | pdf |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.2.6.2. Output {#id59}

Table 2.28. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| pdf | Invoice pdf as base64 encoded | base64 |  |
| raw | Invoice pdf as base64 encoded | text255 |  |


#### 2.2.7. accounting.listInvoices {#accounting.listInvoices}

Get list of available invoices.


##### 2.2.7.1. Input {#id61}

No parameters allowed


##### 2.2.7.2. Output {#id62}

Table 2.29. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of invoices | int |  |
| invoice |  |  |  |
| ... invoiceId | Id of the invoice | text64 |  |
| ... date | Date of invoice | date |  |
| ... afterTax | Bill including tax | float |  |
| ... preTax | Bill without tax | float |  |
| ... type | Kind of invoice | text64 |  |


#### 2.2.8. accounting.lockedFunds {#accounting.lockedFunds}

Log of locked deposit.


##### 2.2.8.1. Input {#id64}

Table 2.30. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| dateFrom | Locked funds start date | timestamp | false |  |
| dateTo | Locked funds end date | timestamp | false |  |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |
| format | Format of the requested document | documentformat | false |  |


##### 2.2.8.2. Output {#id65}

Table 2.31. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| funds |  |  |  |
| ... date | Lock timestamp | dateTime |  |
| ... domain | Domain name for locked fund | text64 |  |
| ... amount | Locked amount value | float |  |
| ... status | Locked domain status | text64 |  |
| sslFunds |  |  |  |
| ... certificateId | Certificate ID | int |  |
| ... creationDate | Date of credit locking | dateTime |  |
| ... commonName | If already set the commonName of the certificate | text255 |  |
| ... netto | Locked certificate status | float |  |
| ... vat | VAT | float |  |
| ... productId | Product ID | int |  |
| ... productName | Name of the product | text255 |  |
| ... status | Locked certificate status | text64 |  |
| pdf | Invoice pdf as base64 encoded | base64 |  |


#### 2.2.9. accounting.log {#accounting.log}

Log of customers transactions.


##### 2.2.9.1. Input {#id67}

Table 2.32. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| dateFrom | Log start date | timestamp | false |  |
| dateTo | Log end date | timestamp | false |  |
| priceMin | Minimum price of log entry | float | false | 0.0 |
| priceMax | Maximum price of log entry | float | false |  |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.2.9.2. Output {#id68}

Table 2.33. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of log entries | int |  |
| sum | Sum of amount | float |  |
| log |  |  |  |
| ... date | Log timestamp | dateTime |  |
| ... amount | Amount value | float |  |
| ... type | Type of action | text64 |  |
| ... details | Log details | text64 |  |
| ... refundable | Indicates whether the payment is refundable or not | boolean |  |


#### 2.2.10. accounting.refund {#accounting.refund}

Request a refund of your unneeded funds.


##### 2.2.10.1. Input {#id70}

Table 2.34. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| creditlogid | ID of the deposit to be refunded | int | true |  |
| amount | Amount to be refunded | float | true |  |


##### 2.2.10.2. Output {#id71}

No additional return parameters


#### 2.2.11. accounting.sendbypost {#accounting.sendbypost}

Requests postal delivery of an invoice.


##### 2.2.11.1. Input {#id73}

Table 2.35. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| invoiceId | Id of the invoice | text64 | true |  |
| acceptCosts | Set to true to accept the incurred costs | boolean | true | false |
| sendTwice | Set to true to accept duplicated delivery-requests for a single invoice | boolean | false | false |
| testing | Testing mode, no real action | boolean | false |  |


##### 2.2.11.2. Output {#id74}

No additional return parameters


## 2.3. Application {#id2904}

The application object provides methods to manage (create, update, delete etc.) domain applications.


#### 2.3.1. application.check {#application.check}

Check availability of domain applications.


##### 2.3.1.1. Input {#id76}

Table 2.36. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | text64 | true |  |


##### 2.3.1.2. Output {#id77}

Table 2.37. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| domain | Domain name | text64 |  |
| price | Domain application price | float |  |
| period | Domain registration period | int |  |
| scDate | Time of scheduled execution | timestamp |  |
| status | Domain check status | text10 |  |
| appCount | Total number of equivalent applications | int |  |


#### 2.3.2. application.create {#application.create}

Create a domain preregistration.


##### 2.3.2.1. Input {#id79}

Table 2.38. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | text64 | true |  |
| price | Domain application price offer | float | false |  |
| registrant | Domain owner contact handle id | int | true |  |
| admin | Domain administrative contact handle id | int | false |  |
| tech | Domain technical contact handle id | int | false |  |
| billing | Domain billing contact handle id | int | false |  |
| ns | List of nameserver | nsList | true |  |
| type | Phase of application | text064 | false |  |
| extData | Domain application extra data | extData | false |  |


##### 2.3.2.2. Output {#id80}

Table 2.39. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain application | int |  |
| notifications | Information about not required contact types | array | Yes |


#### 2.3.3. application.delete {#application.delete}

Delete a domain preregistration.


##### 2.3.3.1. Input {#id82}

Table 2.40. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain application | int | true |  |


##### 2.3.3.2. Output {#id83}

No additional return parameters


#### 2.3.4. application.info {#application.info}

Get domain application details.


##### 2.3.4.1. Input {#id85}

Table 2.41. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain application | int | true |  |
| wide | More detailed output | int | false | 1 |


##### 2.3.4.2. Output {#id86}

Table 2.42. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain application | int |  |
| domain | Domain name of the application | text64 |  |
| domain-ace | Domain ace name of the application | text64 |  |
| type | Phase of application | text64 |  |
| crDate | Time of domain application creation | dateTime | Yes |
| upDate | Time of last domain application update | dateTime | Yes |
| closedDate | Time of domain application close | dateTime | Yes |
| extData | Domain application extra data | extData | Yes |
| price | Domain application price | float | Yes |
| status | Status of application | text64 | Yes |
| registrant | Domain owner contact handle id | int | Yes |
| admin | Domain administrative contact handle id | int | Yes |
| tech | Domain technical contact handle id | int | Yes |
| billing | Domain billing contact handle id | int | Yes |
| ns | List of nameserver | nsList | Yes |
| appPosition | Application queue position | text10 | Yes |
| appCount | Total number of equivalent applications | int | Yes |
| appPrices | Highest application prices | array_float | Yes |
| contact |  |  | Yes |
| ... registrant | Registrant contact handle details | contact |  |
| ... admin | Administrative contact handle details | contact |  |
| ... tech | Technical contact handle details | contact |  |
| ... billing | Billing contact handle details | contact |  |


#### 2.3.5. application.list {#application.list}

List all doamin preregistrations.


##### 2.3.5.1. Input {#id88}

Table 2.43. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain search string | array_text64 | false | * |
| wide | More detailed output | int | false | 0 |
| order | Sort order of result list | applicationOrder | false | DOMAINASC |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.3.5.2. Output {#id89}

Table 2.44. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of preregistrations | int |  |
| sum | Sum of preregistration prices | float |  |
| domain |  |  |  |
| ... roId | Id (Repository Object Identifier) of the domain application | int |  |
| ... domain | Domain name of the application | text64 |  |
| ... domain-ace | Domain ace name of the application | text64 |  |
| ... type | Phase of application | text64 |  |
| ... registrant | Domain owner contact handle id | int | Yes |
| ... admin | Domain administrative contact handle id | int | Yes |
| ... tech | Domain technical contact handle id | int | Yes |
| ... billing | Domain billing contact handle id | int | Yes |
| ... crDate | Time of creation | dateTime | Yes |
| ... upDate | Time of last last update | dateTime | Yes |
| ... closedDate | Time of application close | dateTime | Yes |
| ... extData | Domain application extra data | extData | Yes |
| ... price | Domain application price | float | Yes |
| ... status | Status of application | text64 | Yes |
| ... ns | List of nameserver | nsList | Yes |
| ... appPosition | Application queue position | text10 | Yes |
| ... appCount | Total number of equivalent applications | int | Yes |
| ... appPrices | Highest application prices | array_float | Yes |
| ... contact |  |  | Yes |
| ... ... registrant | Registrant contact handle details | contact |  |
| ... ... admin | Administrative contact handle details | contact |  |
| ... ... tech | Technical contact handle details | contact |  |
| ... ... billing | Billing contact handle details | contact |  |


#### 2.3.6. application.update {#application.update}

Update domain preregistration.


##### 2.3.6.1. Input {#id91}

Table 2.45. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain application | int | true |  |
| price | Domain application price offer | float | false |  |
| registrant | Domain owner contact handle id | int | false |  |
| admin | Domain administrative contact handle id | int | false |  |
| tech | Domain technical contact handle id | int | false |  |
| billing | Domain billing contact handle id | int | false |  |
| ns | List of nameserver | nsList | false |  |
| type | Phase of application | text064 | false |  |
| extData | Domain application extra data | extData | false |  |


##### 2.3.6.2. Output {#id92}

Table 2.46. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| notifications | Information about not required contact types | array | Yes |


## 2.4. Authinfo2 {#id3639}

The authinfo2 object provides methods to get prices and create a new authinfo2 code for .de domains.


#### 2.4.1. authinfo2.create {#authinfo2.create}

Create a new AuthInfo2 product


##### 2.4.1.1. Input {#id94}

Table 2.47. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | AuthInfo2 Domain | text255 | true |  |
| testing | Execute order in testing mode (no order will be submitted) | boolean | false | false |


##### 2.4.1.2. Output {#id95}

Table 2.48. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| product | The name of the bought product | text255 |  |
| net_price | The net price of the bought AuthInfo2 product | float |  |
| price | The price of the bought AuthInfo2 product | float |  |
| currency | Currency of the price | text255 |  |
| message | Message details of the transaction | text255 |  |


#### 2.4.2. authinfo2.getprice {#authinfo2.getprice}

Get the price of a AuthInfo2 product


##### 2.4.2.1. Input {#id97}

No parameters allowed


##### 2.4.2.2. Output {#id98}

Table 2.49. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| product | The name of the product | text255 |  |
| net_price | The net price of the AuthInfo2 product | float |  |
| price | The price of the AuthInfo2 product | float |  |
| currency | Currency of the price | text255 |  |


## 2.5. Certificate {#id3766}

The certificate object provides methods to manage (create, update, cancel etc.) your SSL certificates.


#### 2.5.1. certificate.cancel {#certificate.cancel}

Cancel pending certificate requests


##### 2.5.1.1. Input {#id100}

Table 2.50. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate | int | true |  |


##### 2.5.1.2. Output {#id101}

No additional return parameters


#### 2.5.2. certificate.create {#certificate.create}

Order a certificate


##### 2.5.2.1. Input {#id103}

Table 2.51. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| productId | ID of the product to be purchased | int | true |  |
| csr | Certificate signing request (CSR) | text | false |  |
| commonName | Main domain of the certificate | text255 | false |  |
| period | Term of the certificate (1 or 2 years) | int | false | 1 |
| validationMethod | Method to approve, that the user has control over the domain, which is specified in the CSR. It can be EMAIL, DNS or FILE. | text064 | false |  |
| validationEmail | Email address required to verify, that the user has control over the domain, when EMAIL is chosen as validationMethod | email | false |  |
| ownerc | Contact ID of the owner of the certificate | int | false |  |
| adminc | Contact ID to use as admin | int | false |  |
| techc | Contact ID to use as tech | int | false |  |
| numberOfSan | Maximum number of SAN to be purchased (if no value is specified, freeSanIncluded of product will be set) | int | false | 0 |
| hosting | Indicates, if the certificate is bought together with a hosting package | boolean | false | false |
| autorenew | Should autorenew be activated? | boolean | false |  |
| testing | Execute order in testing mode (no order will be submitted) | boolean | false | false |


##### 2.5.2.2. Output {#id104}

Table 2.52. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| certificateId | ID of the certificate | int |  |
| status | Status of the certificate | text255 |  |
| price | Total price of the certificate with SAN and included VAT | int |  |
| missing | The remaining values ​​needed to complete the order. Use the updateorder function to deliver the remaining values | array_text |  |


#### 2.5.3. certificate.getProduct {#certificate.getProduct}

Get info about a product


##### 2.5.3.1. Input {#id106}

Table 2.53. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| productId | ID of the product | int | true |  |


##### 2.5.3.2. Output {#id107}

Table 2.54. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the product | int |  |
| name | Name of the product | text255 |  |
| brand | Brand of the product | text255 |  |
| description | Product description | text255 |  |
| type | Certificate type (DV, OV, EV) | text255 |  |
| san | Number of SAN, that can be purchased with this product | int |  |
| freeSanIncluded | Number of free SAN | int |  |
| wildcard | Indicates if a product is a wildcard certificate | boolean |  |
| issuanceTime | Issuance time until certificate is available | text255 |  |
| issuanceTimeUnit | Time unit of the issuance time | text255 |  |
| warranty | This is the highest possible amount of financial reimbursement to a client when a certificate is issued to an unauthorised party, leading to financial loss for the client | int |  |
| warrantyCurrency | Currency of the warranty | text255 |  |
| freeReissue | Indicates whether a reissue of the certificate is free | boolean |  |
| price1year | Price of the certificate for a 1 year term (without VAT) | float |  |
| currency | Currency of the prices | text255 |  |
| pricePerSan1Year | Price per SAN for a 1 year term | float |  |


#### 2.5.4. certificate.info {#certificate.info}

Info about a purchased certificate


##### 2.5.4.1. Input {#id109}

Table 2.55. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate | int | true |  |


##### 2.5.4.2. Output {#id110}

Table 2.56. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the certificate | int |  |
| productId | ID of the product | int |  |
| creationDate | Date of the creation of the order | dateTime |  |
| expirationDate | Expiration date of the certificate | dateTime |  |
| updatedAt | Date of the last update of the certificate | dateTime |  |
| period | Term of the certificate (1 or 2 years) | int |  |
| csr | Certificate signing request (CSR) | text |  |
| certificate | SSL Certificate | text |  |
| caChain | Root and intermediate certificates of the CA | text |  |
| validationMethod | Method to approve, that the user has control over the domain, which is specified in the CSR. It can be EMAIL, DNS or FILE. | text064 |  |
| validationEmail | Email address required to verify, that the user has control over the domain, when EMAIL is chosen as validationMethod | email | Yes |
| validationToken | Token used to verify, that the user has control over the domain, when DNS or FILE is chosen as validationMethod | text0 | Yes |
| ownerc | Contact ID of the owner of the certificate | int |  |
| adminc | Contact ID to use as admin | int |  |
| techc | Contact ID to use as tech | int |  |
| status | Status of the certificate | text255 |  |
| commonName | Main domain of the certificate specified in CSR | array_text |  |
| numberOfSan | Maximum number of SAN (specified in CSR) you can use | int |  |
| active | Indicates if the certificate is usable | boolean |  |
| name | Name of the product | text255 |  |
| brand | Brand of the product | text255 |  |
| type | Certificate type (DV, OV, EV) | text255 |  |
| renewedFrom | ID of the predecessor certificate | int |  |
| hosting | Indicates if hosting option is used | boolean |  |
| autorenew | Indicates if certificate is renewed automatically | boolean |  |
| tracenumber | Order reference | text |  |
| san | SAN included | array_text |  |


#### 2.5.5. certificate.list {#certificate.list}

Get a list of all your purchased certificates


##### 2.5.5.1. Input {#id112}

Table 2.57. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| showInactive | Should inactive certificates be listed too? | boolean | false | false |
| ownerc | Filter by owner handle ID | text255 | false |  |
| adminc | Filter by admin handle ID | text255 | false |  |
| techc | Filter by tech handle ID | text255 | false |  |
| validationEmail | Filter by validationEmail | text255 | false |  |
| status | Filter by status | text255 | false |  |
| commonName | Filter by commonName | text255 | false |  |
| sortColumn | Sort result according to a specified column | text255 | false |  |
| order | Sort the specified column in ascending (ASC) or descending (DESC) order | text255 | false |  |


##### 2.5.5.2. Output {#id113}

Table 2.58. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the certificate | int |  |
| productId | ID of the product | int |  |
| creationDate | Date of the creation of the order | dateTime |  |
| expirationDate | Expiration date of the certificate | dateTime |  |
| updatedAt | Date of the last update of the certificate | dateTime |  |
| csr | Certificate signing request (CSR) | text |  |
| certificate | SSL Certificate | text |  |
| caChain | Root and intermediate certificates of the CA | text |  |
| validationMethod | Method to approve, that the user has control over the domain, which is specified in the CSR. It can be EMAIL, DNS or FILE. | text064 |  |
| validationEmail | Email address required to verify, that the user has control over the domain, when EMAIL is chosen as validationMethod | email | Yes |
| validationToken | Token used to verify, that the user has control over the domain, when DNS or FILE is chosen as validationMethod | text0 | Yes |
| ownerc | Contact ID of the owner of the certificate | int |  |
| adminc | Contact ID to use as admin | int |  |
| techc | Contact ID to use as tech | int |  |
| status | Status of the certificate | text255 |  |
| commonName | Main domain of the certificate specified in CSR | array_text |  |
| numberOfSan | Maximum number of SAN (specified in CSR) you can use | int |  |
| active | Indicates if the certificate is usable | boolean |  |
| hosting | Indicates if certificate is used for hosting | boolean |  |
| autorenew | Indicates if the certificate will be renewed automatically | boolean |  |
| name | Name of the product | text255 |  |
| brand | Brand of the product | text255 |  |
| type | Certificate type (DV, OV, EV) | text255 |  |


#### 2.5.6. certificate.listProducts {#certificate.listProducts}

Get a list of all certificate-products with prices


##### 2.5.6.1. Input {#id115}

Table 2.59. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| name | Name of the product | text255 | false |  |
| brand | Brand of the product | text255 | false |  |
| type | Certificate type (DV, OV, EV) | text255 | false |  |
| wildcard | Indicates if a product is a wildcard certificate | boolean | false |  |
| sortColumn | Sort result according to a specified column | text255 | false |  |
| order | Sort the specified column in ascending (ASC) or descending (DESC) order | text255 | false |  |


##### 2.5.6.2. Output {#id116}

Table 2.60. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the product | int |  |
| name | Name of the product | text255 |  |
| brand | Brand of the product | text255 |  |
| description | Product description | text255 |  |
| type | Certificate type (DV, OV, EV) | text255 |  |
| san | Number of SAN, that can be purchased with this product | int |  |
| freeSanIncluded | Number of free SAN | int |  |
| wildcard | Indicates if a product is a wildcard certificate | boolean |  |
| issuanceTime | Issuance time until certificate is available | text255 |  |
| issuanceTimeUnit | Time unit of the issuance time | text255 |  |
| warranty | This is the highest possible amount of financial reimbursement to a client when a certificate is issued to an unauthorised party, leading to financial loss for the client | int |  |
| warrantyCurrency | Currency of the warranty | text255 |  |
| freeReissue | Indicates whether a reissue of the certificate is free | boolean |  |
| price1year | Price of the certificate for a 1 year term (without VAT) | float |  |
| currency | Currency of the prices | text255 |  |
| pricePerSan1Year | Price per SAN for a 1 year term | float |  |


#### 2.5.7. certificate.listRemainingNeededData {#certificate.listRemainingNeededData}

List missing data for completing a specified request


##### 2.5.7.1. Input {#id118}

Table 2.61. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate | int | true |  |


##### 2.5.7.2. Output {#id119}

Table 2.62. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| certificateId | ID of the certificate | text |  |
| missing | Array of missing values for requesting certificate (csr, validationMethod, validationEmail, ownerc, adminc, techc) | array_text |  |


#### 2.5.8. certificate.log {#certificate.log}

Request log of all certificates or of a specific certificate.


##### 2.5.8.1. Input {#id121}

Table 2.63. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate | int | false |  |


##### 2.5.8.2. Output {#id122}

Table 2.64. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| certificateId | ID of the certificate | int |  |
| commonName | Main domain of the certificate specified in CSR | array_text |  |
| creationDate | Date of the creation of log entry | dateTime |  |
| status | Status of the certificate | text255 |  |
| price | Total price of the certificate with SAN and included VAT | float |  |
| netto | Total price of the certificate with SAN and no included VAT | float |  |
| vat | Your Value added tax percentage | float |  |
| statusDetails | Details of certificate status | text1024 |  |
| name | Name of the product | text255 |  |
| brand | Brand of the product | text255 |  |


#### 2.5.9. certificate.reissue {#certificate.reissue}

Reissue a certificate


##### 2.5.9.1. Input {#id124}

Table 2.65. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate to be reissued | int | true |  |
| csr | New certificate signing request (CSR) | text | true |  |


##### 2.5.9.2. Output {#id125}

Table 2.66. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| certificateId | ID of the certificate | int |  |
| status | Status of the certificate | text255 |  |


#### 2.5.10. certificate.renew {#certificate.renew}

Renew a certificate


##### 2.5.10.1. Input {#id127}

Table 2.67. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate to be renewed | int | true |  |
| csr | Certificate signing request (CSR) | text | false |  |
| commonName | Main domain of the certificate | text255 | false |  |
| validationMethod | Method to approve, that the user has control over the domain, which is specified in the CSR. It can be EMAIL, DNS or FILE. | text064 | false |  |
| validationEmail | Email address required to verify, that the user has control over the domain, when EMAIL is chosen as validationMethod | email | false |  |
| ownerc | Contact ID of the owner of the certificate | int | false |  |
| adminc | Contact ID to use as admin | int | false |  |
| techc | Contact ID to use as tech | int | false |  |
| numberOfSan | Maximum number of SAN to be purchased (if no value is specified, freeSanIncluded of product will be set) | int | false | 0 |
| autorenew | Should autorenew be activated? | boolean | false |  |
| testing | Execute order in testing mode (no order will be submitted) | boolean | false | false |


##### 2.5.10.2. Output {#id128}

Table 2.68. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| certificateId | ID of the new certificate | int |  |
| status | Status of the new certificate | text255 |  |
| price | Total price of the new certificate with SAN and included VAT | int |  |


#### 2.5.11. certificate.setAutorenew {#certificate.setAutorenew}

Change the autorenew setting for a certificate


##### 2.5.11.1. Input {#id130}

Table 2.69. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate | int | true |  |
| autorenew | Should autorenew be activated? | boolean | true |  |


##### 2.5.11.2. Output {#id131}

No additional return parameters


#### 2.5.12. certificate.setresendapproval {#certificate.setresendapproval}

Change the resend approval mail setting for a certificate


##### 2.5.12.1. Input {#id133}

Table 2.70. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate | int | true |  |
| resend_approval | Resend the mail for approving the certificate? | boolean | true |  |


##### 2.5.12.2. Output {#id134}

No additional return parameters


#### 2.5.13. certificate.updateOrder {#certificate.updateOrder}

Submit missing information of your certificate request.


##### 2.5.13.1. Input {#id136}

Table 2.71. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| certificateId | ID of the certificate | int | true |  |
| csr | Certificate signing request (CSR) | text | false |  |
| commonName | Main domain of the certificate | text255 | false |  |
| validationMethod | Method to approve, that the user has control over the domain, which is specified in the CSR. It can be EMAIL, DNS or FILE. | text064 | false |  |
| validationEmail | Email address required to verify, that the user has control over the domain, when EMAIL is chosen as validationMethod | email | false |  |
| ownerc | Contact ID of the owner of the certificate | int | false |  |
| adminc | Contact ID to use as admin | int | false |  |
| techc | Contact ID to use as tech | int | false |  |
| hosting | Indicates, if the certificate is bought together with a hosting package | boolean | false | false |
| autorenew | Should autorenew be activated? | boolean | false |  |


##### 2.5.13.2. Output {#id137}

Table 2.72. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| certificateId | ID of the certificate | int |  |
| status | Status of the certificate | text255 |  |
| missing | The remaining values ​​needed to complete the order. Use the updateorder function to deliver the remaining values | array_text |  |


## 2.6. Contact {#id5242}

The contact object provides methods to manage (create, update, delete etc.) your contact handles.


#### 2.6.1. contact.create {#contact.create}

Creates a new contact handle.


##### 2.6.1.1. Input {#id139}

Table 2.73. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| type | Type of contact | contactType | true |  |
| name | First and lastname | text255 | true |  |
| org | Organisation/company | text0255 | false |  |
| street | Street address field 1 | text64 | true |  |
| city | Contact city | text64 | true |  |
| pc | Contact postal code | text10 | true |  |
| sp | Contact state province | text064 | false |  |
| cc | Contact country | country | true |  |
| voice | Contact phone number | phone | true |  |
| fax | Contact fax number | phone | false |  |
| email | Contact email address | email | true |  |
| remarks | Contact handle remarks | remarks | false |  |
| forceNew | Force new contact handle creation | boolean | false | 0 |
| extData | Contact extra data | extData | false |  |
| testing | Execute command in testing mode | boolean | false | 0 |


##### 2.6.1.2. Output {#id140}

Table 2.74. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | Contact handle id | int |  |


#### 2.6.2. contact.delete {#contact.delete}

Delete an existing contact handle.


##### 2.6.2.1. Input {#id142}

Table 2.75. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Contact handle id | int | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.6.2.2. Output {#id143}

No additional return parameters


#### 2.6.3. contact.info {#contact.info}

Get contact handle details.


##### 2.6.3.1. Input {#id145}

Table 2.76. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Contact handle id | int | true |  |
| wide | More detailed output | int | false | 0 |


##### 2.6.3.2. Output {#id146}

Table 2.77. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| contact |  |  |  |
| ... roId | Contact handle id | int |  |
| ... id | Contact handle id name | int |  |
| ... type | Type of contact handle | contactType |  |
| ... name | First and lastname | text255 |  |
| ... org | Organisation/company | text0255 |  |
| ... street | Street address | text255 |  |
| ... city | City | text64 |  |
| ... pc | Postal code | text10 |  |
| ... sp | State province | text064 | Yes |
| ... cc | Country | country |  |
| ... voice | Phone number | phone |  |
| ... fax | Fax number | phone | Yes |
| ... email | Email address | email |  |
| ... remarks | Contact handle remarks | remarks | Yes |
| ... usedCount | Total number of contact handle uses | int | Yes |
| ... nicHandle | Contact NIC handle | nicHandleList | Yes |


#### 2.6.4. contact.list {#contact.list}

List available contact handles.


##### 2.6.4.1. Input {#id148}

Table 2.78. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| order | Sort order of the result list | contactOrder | false | IDDESC |
| search | Search string | text64 | false |  |
| readOnly | List only readable contact handle | int | false |  |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |
| withoutVerification | Don't return the Verification Status | boolean | false | 0 |
| id | Retrieve only this one contact. | int | false |  |


##### 2.6.4.2. Output {#id149}

Table 2.79. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of contact handles | int |  |
| contact |  |  |  |
| ... roId | Contact handle id (Repository Object Identifier) | int |  |
| ... id | Contact handle id name | int |  |
| ... type | Type of contact handle | contactType |  |
| ... name | First and lastname | text255 |  |
| ... org | Organisation/company | text0255 |  |
| ... street | Street address | text255 |  |
| ... city | City | text64 |  |
| ... pc | Postal code | text10 |  |
| ... sp | State province | text064 | Yes |
| ... cc | Country | country |  |
| ... voice | Phone number | phone |  |
| ... fax | Fax number | phone | Yes |
| ... email | Email adress | email |  |
| ... remarks | Contact handle remarks | remarks | Yes |
| ... readOnly | Contact handle is read only | boolean |  |
| ... usedCount | Total number of contact handle uses | int | Yes |
| ... verificationStatus | Contact Verification Status | text64 | Yes |


#### 2.6.5. contact.log {#contact.log}

Log of changes to the contact handle.


##### 2.6.5.1. Input {#id151}

Table 2.80. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Contact handle id | int | false |  |


##### 2.6.5.2. Output {#id152}

Table 2.81. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of log entries | int |  |
| contact |  |  |  |
| ... logId | Id of the log entry | int |  |
| ... date | Log timestamp | dateTime |  |
| ... id | Id of the altered contact handle | int |  |
| ... status | Contact handle status after the action was performed | text64 |  |


#### 2.6.6. contact.sendbulkverification {#contact.sendbulkverification}

send bulk Contact Verification


##### 2.6.6.1. Input {#id154}

Table 2.82. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| search | Search term to filter contacts | string | false |  |
| verificationStatus | Array of verification statuses to filter contacts (CONFIRMED, AWAIT_CONFIRMATION, NONE, TO_NOTIFY) | array | false |  |


##### 2.6.6.2. Output {#id155}

Table 2.83. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| code | Return code (1000 = success) | int |  |
| msg | Error message if code != 1000 | string |  |


#### 2.6.7. contact.sendcontactverification {#contact.sendcontactverification}

send Contact Verification


##### 2.6.7.1. Input {#id157}

Table 2.84. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Contact handle id | int | true |  |


##### 2.6.7.2. Output {#id158}

No additional return parameters


#### 2.6.8. contact.update {#contact.update}

Alter contact handle data.


##### 2.6.8.1. Input {#id160}

Table 2.85. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Contact handle id | int | true |  |
| name | First and lastname | text255 | false |  |
| org | Organisation/company | text0255 | false |  |
| street | Street address field 1 | text64 | false |  |
| city | City | text64 | false |  |
| pc | Postal code | text10 | false |  |
| sp | State province | text064 | false |  |
| cc | Country | country | false |  |
| voice | Contact phone number | text64 | false |  |
| fax | Contact fax number | text064 | false |  |
| email | Email address | email | false |  |
| remarks | Contact handle remarks | remarks | false |  |
| extData | Contact extra data | extData | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.6.8.2. Output {#id161}

No additional return parameters


## 2.7. Customer {#id6062}

The customer object provides methods to manage common properties for all your accounts.


#### 2.7.1. customer.contactverificationsettingsinfo {#customer.contactverificationsettingsinfo}

Get contact verification settings.


##### 2.7.1.1. Input {#id163}

No parameters allowed


##### 2.7.1.2. Output {#id164}

Table 2.86. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| supportEmail | Email address. | email |  |
| supportName | Email name. | text255 |  |
| signature | Signature. | text0 |  |


#### 2.7.2. customer.contactverificationsettingsupdate {#customer.contactverificationsettingsupdate}

Update contact verification email settings.


##### 2.7.2.1. Input {#id166}

Table 2.87. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| supportEmail | Customer supporter's email address. | emailoptional | false |  |
| supportName | Customer supporter's name. | text0255 | false |  |
| signature | Customer specific signature. | text0 | false |  |


##### 2.7.2.2. Output {#id167}

No additional return parameters


#### 2.7.3. customer.delete {#customer.delete}

Delete a customer.


##### 2.7.3.1. Input {#id169}

No parameters allowed


##### 2.7.3.2. Output {#id170}

No additional return parameters


#### 2.7.4. customer.info {#customer.info}

Get customer information.


##### 2.7.4.1. Input {#id172}

Table 2.88. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| wide | More detailed output | int | false |  |


##### 2.7.4.2. Output {#id173}

Table 2.89. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| accountId | Account id | int |  |
| customerId | Customer id | int |  |
| customerNo | Your unique Customer number | int |  |
| username | Username | text64 |  |
| title | Salutation title | addressTitle |  |
| firstname | Customers firstname | text64 |  |
| lastname | Customers lastname | text64 |  |
| org | Customers organisation/company | text064 |  |
| street | Customers street | text255 |  |
| pc | Customers postal code | text10 |  |
| city | Customers city | text64 |  |
| cc | Customers country | country |  |
| voice | Customers phone number | phone |  |
| fax | Customers fax number | phoneoptional |  |
| www | Customers web address | token0255 |  |
| email | Customers email address | email |  |
| servicePin | Customers service pin | int |  |
| crDate | Customers date of account creation | dateTime |  |
| secureMode | Customers secure mode | boolean |  |
| signPdfs | Customer gets signed pdfs | boolean |  |
| summaryInvoice | Customer gets summary invoice | boolean |  |
| mailListId | List of subscribed mail-list-ids | array_int | Yes |
| language | Customers language | language |  |
| notificationEmail | Customer gets notification emails | boolean |  |
| notificationQueue | Customer uses notification queue | boolean |  |
| lowBalance | Customers low balance notification value | float |  |
| renewalReport | Customer gets renewal reports | boolean |  |
| paymentType | Customers type of payment | paymentType |  |
| bankAccHolder | Customers name of bank account holder | token255 |  |
| bankName | Customers name of bank | token255 |  |
| bankCode | Customers bank code | token255 |  |
| bankAccHolderNo | Customers bank account number | text64 |  |
| vat | Customers vat | int |  |
| vatNo | Customers company vat number | vatNo |  |
| whoisProvider | Default whois provider | token255 |  |
| whoisUrl | Default whois url | token255 |  |
| defaultRegistrant | Default registrant contact handle | int |  |
| defaultAdmin | Default administrative contact handle | int |  |
| defaultTech | Default technical contact handle | int |  |
| defaultBilling | Default billing contact handle | int |  |
| invoiceXml | Xml type invoices will be send | boolean |  |
| invoicePdf | Pdf type invoices will be send | boolean |  |
| defaultNsset | Default nameserver set | int |  |
| defaultWeb | Default web nameserver entry | token0255 | Yes |
| defaultMail | Default mail nameserver entry | text0255 | Yes |
| defaultImportNS | Import NS in case of Domain Transfer | boolean | Yes |
| defaultRenewalMode | Default domain renewal mode | renewalMode |  |
| lastLogin | Date of last login | dateTime |  |
| loginCount | Count of logins | int |  |
| rowsPerPage | Default rows per page value | int |  |
| tfa | 2-factor-authentification method | tfaMethod |  |
| lastIP | Customers ip address of last access | ip |  |
| verification | Customers data verfication pending flag | int |  |
| emailBilling | Customers email address for billing | email |  |
| emailAutomated | Email address for automated emails | email |  |
| currency | currency | customercurrency |  |
| isReseller | Defines if the customer is a reseller or not | boolean_3 |  |
| serviceProviderId | The provider-id of the customer | int |  |
| supplimentInvoiceText | Suppliment text to the invoice | text0100 |  |
| disablepremium | Are premium domains disabled by the customer? | boolean |  |
| enableAdvertising | Advertising is allowed on expired domains. | boolean |  |
| wdrpEmail | Email address for WDRP notifications | email |  |


#### 2.7.5. customer.listdownloads {#customer.listdownloads}

List downloadable file tokens for the given customer.


##### 2.7.5.1. Input {#id175}

Table 2.90. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| type | Type of download to search for, e.g. 'gdpr'. If omitted, all downloads are returned. | text64 | false |  |


##### 2.7.5.2. Output {#id176}

Table 2.91. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| downloads |  |  |  |
| ... token | Download token | text64 |  |
| ... created | Creation date of this download | dateTime |  |
| ... downloads | Amount of times this file has been downloaded | int |  |
| ... expires | Expire date of this download | dateTime |  |
| ... status | Status of this download | text64 |  |


#### 2.7.6. customer.requestdataexport {#customer.requestdataexport}

Request a customer data export.


##### 2.7.6.1. Input {#id178}

No parameters allowed


##### 2.7.6.2. Output {#id179}

No additional return parameters


#### 2.7.7. customer.update {#customer.update}

Update customer information.


##### 2.7.7.1. Input {#id181}

Table 2.92. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| username | Customers username | username | false |  |
| title | Customers salutation title | addressTitle | false |  |
| firstname | Customers firstname | text64 | false |  |
| lastname | Customers lastname | text64 | false |  |
| org | Customers organisation/company | token0255 | false |  |
| street | Customers street | text255 | false |  |
| pc | Customers postal code | text10 | false |  |
| city | Customers city | text64 | false |  |
| cc | Customers country | country | false |  |
| voice | Customers phone number | phone | false |  |
| fax | Customers fax number | phoneoptional | false |  |
| www | Customers web address | token0255 | false |  |
| email | Customers email address | email | false |  |
| emailAutomated | Customers automated email address | emailoptional | false |  |
| emailBilling | Customers email address for billing | emailoptional | false |  |
| servicePin | Customers service pin | text64 | false |  |
| summaryInvoice | Customer gets summary invoice | boolean | false |  |
| mailListId | One or more mail-list-id(s) | array_int | false |  |
| language | Customers language | language | false |  |
| notificationEmail | Customer get notification emails | boolean | false |  |
| notificationQueue | Customer uses notification queue | boolean | false |  |
| renewalReport | Customer gets renewal reports emails | boolean | false |  |
| lowBalance | Customers low balance notification value | float_signed | false |  |
| bankAccHolder | Customers name of bank account holder | token255 | false |  |
| bankCode | Customers bank code | text64 | false |  |
| bankName | Customers name of bank | text64 | false |  |
| bankAccHolderNo | Customers bank account number | text64 | false |  |
| vatNo | Customers company vat number | vatNoInternational | false |  |
| whoisProvider | Default whois provider | token0255 | false |  |
| whoisUrl | Default whois url | token0255 | false |  |
| defaultRegistrant | Default registrant contact handle | int | false |  |
| defaultAdmin | Default admin contact handle | int | false |  |
| defaultTech | Default tech contact handle | int | false |  |
| defaultBilling | Default billing contact handle | int | false |  |
| invoiceXml | Xml type invoices will be send | boolean | false |  |
| invoicePdf | Pdf type invoices will be send | boolean | false |  |
| defaultNsset | Default nameserver set | int | false |  |
| defaultWeb | Default web nameserver entry | token0255 | false |  |
| defaultMail | Default mail nameserver entry | text0255 | false |  |
| defaultImportNS | Import NS in Case of Domain Transfer | boolean | false |  |
| defaultRenewalMode | Default domain renewal mode | renewalMode | false |  |
| supplimentinvoicetext | Suppliment text to the invoice | text0100 | false |  |
| password | Customers password | password | false |  |
| rowsPerPage | Default rows per page value | int | false |  |
| isReseller | Is Customer a reseller? | boolean_3 | false |  |
| testing | Execute command in testing mode | boolean | false | false |
| nif | NIF of a natural person or company (only for customers of InternetworX SLU) | text10 | false |  |
| idCard | Your ID card number (only for natural person with residence in the EU for InternetworX SLU) | text64 | false |  |
| disablepremium | Are premium domains disabled by the customer? | boolean | false |  |
| enableAdvertising | Advertising is allowed on expired domains. | boolean | false |  |
| wdrpEmail | Specify the email address for WDRP notifications | emailoptional | false |  |
| errpEmail | Specify the sender email address for ERRP notifications | emailoptional | false |  |
| errpEmailName | Specify the sender name for ERRP notifications | text0255 | false |  |
| renewalReportEmail | Specify the email address for renewal report notifications | emailoptional | false |  |


##### 2.7.7.2. Output {#id182}

No additional return parameters


## 2.8. Dnssec {#id7052}

The dnssec object provides methods to manage DNSSEC for your domains.


#### 2.8.1. dnssec.adddnskey {#dnssec.adddnskey}

Add one DNSKEY to a specified domain. Currently, only ZONE+SEP keys (flag value 257) are accepted. This does not overwrite or delete existing DNSKEYs to allow for key rollovers.


##### 2.8.1.1. Input {#id184}

Table 2.93. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domainName | Name of the domain to add the DNSKEY for. | text255 | true |  |
| dnskey | Presentation value for the DNSKEY to add. Example: domain.tld. IN DNSKEY 257 3 13 ac12c2... | dnskey | false |  |
| ds | Optional presentation value for the corresponding DS record (digest information). Example: domain.tld. IN DS 1234 13 2 56DC12... | ds | false |  |
| calculateDigest | If TRUE, the digest values for this DNSKEY will be calculated. Overrides ds parameter. | boolean | false | false |
| digestType | This value determines the type of digest which will be calculated. Defaults to 2 (SHA256). | dnssecDigestType | false | 2 |


##### 2.8.1.2. Output {#id185}

Table 2.94. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| dnskey | Presentation value for the DNSKEY. Example: domain.tld. IN DNSKEY 257 3 13 ac12c2... | dnskey |  |
| ds | Optional presentation value for the corresponding DS record (digest information). Example: domain.tld. IN DS 1234 13 2 56DC12... | ds |  |


#### 2.8.2. dnssec.deleteall {#dnssec.deleteall}

Delete all DNSKEY/DS entries for a domain.


##### 2.8.2.1. Input {#id187}

Table 2.95. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domainName | Name of the domain to delete all DNSKEY/DS records for. | text255 | true |  |


##### 2.8.2.2. Output {#id188}

No additional return parameters


#### 2.8.3. dnssec.deletednskey {#dnssec.deletednskey}

Delete one DNSKEY from a specified domain.


##### 2.8.3.1. Input {#id190}

Table 2.96. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| key | ID of the DNSKEY to delete. | int | true |  |


##### 2.8.3.2. Output {#id191}

No additional return parameters


#### 2.8.4. dnssec.disablednssec {#dnssec.disablednssec}

Disable automated DNSSEC management for a domain. This flags the domain for DNSKEY removal - all keys will be destroyed.


##### 2.8.4.1. Input {#id193}

Table 2.97. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domainName | Name of the domain to disable DNSSEC for. | text255 | true |  |


##### 2.8.4.2. Output {#id194}

No additional return parameters


#### 2.8.5. dnssec.enablednssec {#dnssec.enablednssec}

Enable automated DNSSEC management for a domain.


##### 2.8.5.1. Input {#id196}

Table 2.98. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domainName | Name of the domain to enable DNSSEC for. | text255 | true |  |


##### 2.8.5.2. Output {#id197}

No additional return parameters


#### 2.8.6. dnssec.info {#dnssec.info}

Get current DNSSEC information.


##### 2.8.6.1. Input {#id199}

Table 2.99. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domains | Optionally limit info to given domains. | array_text255 | false |  |


##### 2.8.6.2. Output {#id200}

Table 2.100. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| data | List of domains and their DNSSEC status | array |  |
| ... domain | Domain name. | text255 |  |
| ... keyCount | Count of DNSSEC keys for this domain. | int |  |
| ... dnssecStatus | Domain DNSSEC status. | dnssecDomainStatus |  |


#### 2.8.7. dnssec.listkeys {#dnssec.listkeys}

Search and list manually managed DNSSEC keys.


##### 2.8.7.1. Input {#id202}

Table 2.101. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domainName | Search for DNSSEC data for the given domain. | text0255 | false |  |
| domainNameIdn | Search for DNSSEC data for the given ACE domain name. | text0255 | false |  |
| keyTag | Search for DNSKEY entries with the given key tag. | int | false |  |
| flagId | Search for DNSKEY entries with the given flags value. | dnssecFlag | false |  |
| algorithmId | Search for DNSKEY entries with the given algorithm. | dnssecAlgorithm | false |  |
| publicKey | Search for DNSKEY entries with the given public key. | text | false |  |
| digestTypeId | Search DNSKEY entries with the given digest type. | dnssecDigestType | false |  |
| digest | Search DNSKEY entries with the given digest. | text0255 | false |  |
| createdBefore | Search DNSKEY entries created before this time. | dateTime | false |  |
| createdAfter | Search DNSKEY entries created after this time. | dateTime | false |  |
| status | Search DNSKEY entries with this status. | dnssecKeyStatus | false |  |
| active | Search DNSKEY entries which are active (1) or inactive (0). | int | false |  |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page. 0 is no limit | int | false | 0 |


##### 2.8.7.2. Output {#id203}

Table 2.102. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| dnskey |  |  |  |
| ... ownerName | The domain name that owns the DNSSEC key. | text0255 |  |
| ... id | The unique identifier for the DNSSEC key. | int |  |
| ... domainId | The identifier for the domain associated with the DNSSEC key. | int |  |
| ... keyTag | The key tag associated with the DNSSEC key. | int |  |
| ... flagId | The flag ID of the DNSSEC key. | dnssecFlag |  |
| ... algorithmId | The algorithm ID used by the DNSSEC key. | dnssecAlgorithm |  |
| ... publicKey | The public key for the DNSSEC key. | text |  |
| ... digestTypeId | The digest type ID associated with the DNSSEC key. | dnssecDigestType |  |
| ... digest | The digest associated with the DNSSEC key. | text0255 |  |
| ... created | The date and time when the DNSSEC key was created. | dateTime |  |
| ... status | The status of the DNSSEC key (e.g., OK, DELETED). | dnssecKeyStatus |  |
| ... active | Indicates if the DNSSEC key is active (1) or inactive (0). | int |  |


## 2.9. Domain {#id7530}

The domain object provides methods for creating, deleting, listing etc. of domains.


#### 2.9.1. domain.check {#domain.check}

Check the availability of domains.


##### 2.9.1.1. Input {#id205}

Table 2.103. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | array_domain | false |  |
| sld | Second level domain name | text64 | false |  |
| tld | Top level domain | array_text64 | false |  |
| region | Check top level domains of the given groups | array_identifier | false |  |
| wide | More detailed output | int | false | 1 |


##### 2.9.1.2. Output {#id206}

Table 2.104. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| domain |  |  |  |
| ... domain | Domain name | token255 |  |
| ... avail | Domain availability for registration | int |  |
| ... status | Domain availability status | text10 | Yes |
| ... reason | Reason of the domain invalid status | text | Yes |
| ... checktime | Domain check time | float | Yes |
| ... name | Second level domain name | text64 | Yes |
| ... tld | Top level domain name | text10 | Yes |
| ... checkmethod | Domain check method | text10 | Yes |
| ... premium | Domain premium prices (e.g. reg, renewal) | array | Yes |
| ... price | Domain registration/transfer price | float | Yes |


#### 2.9.2. domain.create {#domain.create}

Register a domain name.


##### 2.9.2.1. Input {#id208}

Table 2.105. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| period | Domain registration/renewal period | period | false | * |
| registrant | Domain owner contact handle id | int | true |  |
| admin | Domain administrative contact handle id | int | false |  |
| tech | Domain technical contact handle id | int | false |  |
| billing | Domain billing contact handle id | int | false |  |
| ns | List of nameserver | nsList | false |  |
| transferLock | Lock domain | boolean | false | 1 |
| renewalMode | Domain renewal mode | renewalMode | false |  |
| whoisProvider | Whois provider | token0255 | false |  |
| whoisUrl | Whois url | token0255 | false |  |
| scDate | Time of scheduled execution | timestamp | false |  |
| extData | Domain extra data | extData | false |  |
| asynchron | Asynchron domain create | boolean | false | false |
| voucher | Voucher code | text64 | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.2.2. Output {#id209}

Table 2.106. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain | int |  |
| price | Incurred expenses | float |  |
| currency | Currency related to price value | customercurrency |  |
| notifications | Information about not required contact types | array | Yes |


#### 2.9.3. domain.delete {#domain.delete}

Deletes a domain.


##### 2.9.3.1. Input {#id211}

Table 2.107. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| scDate | Time of scheduled execution | timestamp | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.3.2. Output {#id212}

No additional return parameters


#### 2.9.4. domain.getalldomainprices {#domain.getalldomainprices}

Get all domain prices (e.g. reg, renewal) for a specified domain.


##### 2.9.4.1. Input {#id214}

Table 2.108. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | array_domain | true |  |
| period | period for the requested action | text64 | false |  |
| voucher | Voucher code | text64 | false |  |


##### 2.9.4.2. Output {#id215}

Table 2.109. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| domain |  |  |  |
| ... premium | Flag, if the given price is premium | boolean |  |
| ... currency | The currency belonging to the prices | customercurrency |  |
| ... period | The value of the period, that is required for this item, will usually match period from request | text64 |  |
| ... promo | Flag, if the given price is promo | boolean |  |
| ... promoPrice | The Promo price of this item | float |  |
| ... promoType | The type of the perfomed action for this promo action | text64 |  |
| ... regPrice | The registration price of this item | float |  |
| ... updatePrice | The update price of this item | float |  |
| ... transferPrice | The transfer price of this item | float |  |
| ... tradePrice | The trading price of this item | float |  |
| ... renewalPrice | The renewal price of this item | float |  |
| ... restorePrice | The restore price of this item | float |  |
| ... promoPriceWithVat | The Promo price of this item including VAT | float |  |
| ... regPriceWithVat | The registration price of this item including VAT | float |  |
| ... updatePriceWithVat | The update price of this item including VAT | float |  |
| ... transferPriceWithVat | The transfer price of this item including VAT | float |  |
| ... tradePriceWithVat | The trading price of this item including VAT | float |  |
| ... renewalPriceWithVat | The renewal price of this item including VAT | float |  |
| ... restorePriceWithVat | The restore price of this item including VAT | float |  |


#### 2.9.5. domain.getdomainprice {#domain.getdomainprice}

Get the domain price by type (e.g. reg, renewal) for a specified domain.


##### 2.9.5.1. Input {#id217}

Table 2.110. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | array_domain | true |  |
| pricetype | Type to get the price for. (Options: reg | renewal | transfer | update | trade | restore ) | text64 | true |  |
| period | period for the requested action | text64 | false |  |
| voucher | Voucher code | text64 | false |  |


##### 2.9.5.2. Output {#id218}

Table 2.111. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| domain |  |  |  |
| ... type | The type of the perfomed action for this item, most likely the same value as pricetype | text64 |  |
| ... period | The value of the period, that is required for this item, will usually match period from request | text64 |  |
| ... price | The price of this item and action | float |  |
| ... currency | The currency belonging to the price | customercurrency |  |
| ... promo | Flag, if the given price is promo | boolean |  |
| ... premium | Flag, if the given price is premium | boolean |  |


#### 2.9.6. domain.getextradatarules {#domain.getextradatarules}

Get extra rules for one or all TLDs.


##### 2.9.6.1. Input {#id220}

Table 2.112. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| tld | Top level domain | text64 | false |  |


##### 2.9.6.2. Output {#id221}

Table 2.113. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| rules | List with rules and constraints of the top level domains | array_text255 |  |


#### 2.9.7. domain.getPrices {#domain.getPrices}

Get the domain prices.


##### 2.9.7.1. Input {#id223}

Table 2.114. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| tld | Top level domain | array_text64 | false |  |
| vat | Prices with vat | boolean | false | true |
| vatCC | 2-letter ISO country code | country | false |  |
| voucher | Voucher code | text64 | false |  |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page. Default is unlimited | int | false |  |


##### 2.9.7.2. Output {#id224}

Table 2.115. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| price |  |  |  |
| ... tld | Top level domain name | text64 |  |
| ... currency | Currency of the price | text10 |  |
| ... createPrice | Domain registration price/year | float |  |
| ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... transferPrice | Domain transfer price/year | float |  |
| ... renewalPrice | Domain renewal price/year | float |  |
| ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... updatePrice | Domain update price | float |  |
| ... tradePrice | Domain trade price | float |  |
| ... trusteePrice | Domain trustee service price/year | float |  |
| ... monthlyTrusteePrice | Domain trustee service price/month | float | Yes |
| ... createPeriod | Domain creation period in years | int |  |
| ... transferPeriod | Domain transfer period in years | int |  |
| ... renewalPeriod | Domain renewal period in years | int |  |
| ... tradePeriod | Domain trade period in years | int |  |
| ... vat | Added value tax rate according to vat and vatCC settings | float | Yes |
| ... promo |  |  | Yes |
| ... ... currency | Currency of the promo price | text10 |  |
| ... ... limitedNumber | Number of limited promos | int |  |
| ... ... limitedRemaining | Number of remaining promo prices for limited promos | int |  |
| ... ... startTime | Promo start time | dateTime |  |
| ... ... endTime | Promo end time | dateTime |  |
| ... ... period | The period the promo offer applies to | text10 |  |
| ... ... tradePrice | trade promo price | float | Yes |
| ... ... trusteePrice | trustee promo price | float | Yes |
| ... ... createPrice | Domain registration price/year | float | Yes |
| ... ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... ... transferPrice | Domain transfer price/year | float | Yes |
| ... ... renewalPrice | Domain renewal price/year | float | Yes |
| ... ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... ... chargedPrice |  |  | Yes |
| ... ... ... currency | Currency of the promo price | text10 |  |
| ... ... ... createPrice | Domain registration price/year | float |  |
| ... ... ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... ... ... transferPrice | Domain transfer price/year | float |  |
| ... ... ... renewalPrice | Domain renewal price/year | float |  |
| ... ... ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... ... ... updatePrice | Domain update price | float |  |
| ... ... ... tradePrice | Domain trade price | float |  |
| ... ... ... trusteePrice | Domain trustee service price/year | float |  |
| ... ... ... monthlyTrusteePrice | Domain trustee service price/month | float | Yes |
| ... chargedPrice |  |  | Yes |
| ... ... currency | Currency of the promo price | text10 |  |
| ... ... createPrice | Domain registration price/year | float |  |
| ... ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... ... transferPrice | Domain transfer price/year | float |  |
| ... ... renewalPrice | Domain renewal price/year | float |  |
| ... ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... ... updatePrice | Domain update price | float |  |
| ... ... tradePrice | Domain trade price | float |  |
| ... ... trusteePrice | Domain trustee service price/year | float |  |
| ... ... monthlyTrusteePrice | Domain trustee service price/month | float | Yes |
| ... eur |  |  | Yes |
| ... ... createPrice | Domain registration price/year | float |  |
| ... ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... ... transferPrice | Domain transfer price/year | float |  |
| ... ... renewalPrice | Domain renewal price/year | float |  |
| ... ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... ... updatePrice | Domain update price | float |  |
| ... ... tradePrice | Domain trade price | float |  |
| ... ... trusteePrice | Domain trustee service price/year | float |  |
| ... ... monthlyTrusteePrice | Domain trustee service price/month | float | Yes |


#### 2.9.8. domain.getPromos {#domain.getPromos}

Get promo prices of TLDs.


##### 2.9.8.1. Input {#id226}

Table 2.116. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| tlds | List of specific TLDs to check for promos | array_text64 | false |  |
| promoType | Domain action (e.g. REG, RENEWAL) | text64 | false | REG |
| period | Domain period | int | false | 1 |
| periodUnit | Domain period unit (e.g. Y for years or M for months) | text10 | false | Y |
| executionDate | Date of promo request | text64 | false |  |
| voucher | Voucher code | text64 | false |  |


##### 2.9.8.2. Output {#id227}

Table 2.117. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| promos |  |  |  |
| ... tld | TLD name | text64 |  |
| ... netPrice | Net promotional price | float |  |
| ... currency | Currency of net price | customercurrency |  |
| ... endTime | End time of the promo | dateTime |  |


#### 2.9.9. domain.getRules {#domain.getRules}

Get TLD rules.


##### 2.9.9.1. Input {#id229}

Table 2.118. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| tld | Top level domain name | array_text64 | false |  |


##### 2.9.9.2. Output {#id230}

Table 2.119. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| rules |  |  |  |
| ... tld | Top level domain name | text64 |  |
| ... cc | TLD country | country |  |
| ... region | All TLD groups this domain is part of | array_identifier |  |
| ... timeRegistration | Time for registration | text64 |  |
| ... localOwner | Local domain owner contact needed | boolean |  |
| ... localAdmin | Local domain administrative contact needed | boolean |  |
| ... localTech | Local domain technical contact needed | boolean |  |
| ... ownerType | Allowed domain owner contact types | text64 |  |
| ... adminType | Allowed domain administrative contact types | text64 |  |
| ... techType | Allowed domain technical contact types | text64 |  |
| ... billingType | Allowed domain billing contact types | text64 |  |
| ... billingFaxNeeded | Domain billing contact needs a fax number | boolean |  |
| ... companyAllowed | Domain owner contact can be a company | boolean |  |
| ... personAllowed | Domain owner contact can be a natural person | boolean |  |
| ... companyNeeded | Domain owner contact has to be a company | boolean |  |
| ... personNeeded | Domain owner contact has to be a natural person | boolean |  |
| ... VATNeeded | Vat number is needed | boolean |  |
| ... IDNeeded | Identification is needed | boolean |  |
| ... markNeeded | Trademark is needed | boolean |  |
| ... idn | Domain name can contain idn chars | boolean |  |
| ... idnChars | Allowed domain name idn characters | token255 |  |
| ... minLength | Minimal domain name length | int |  |
| ... maxLength | Maximal domain name length | int |  |
| ... minNS | Minimum amount of nameservers | int |  |
| ... maxNS | Maximum amount of nameservers | int |  |
| ... dnssec | Which type of DNSSEC record type is supported | text64 |  |
| ... registrationPeriod | Allowed domain registration periods | text64 |  |
| ... renewalPeriod | Allowed domain renewal periods | text64 |  |
| ... restorePeriod1 | Number of days after a EXPIRE INITIATED event in which a restore is possible | int |  |
| ... restorePeriod2 | Number of days after a DELETE INITIATED event in which a restore is possible | int |  |
| ... trustee | Trustee service available | boolean |  |
| ... privacyLevel | Jurisdiction provides data protection for whois data | text |  |
| ... whoisExposure | Personal information is publicly available through whois | text |  |
| ... whoisProtection | Whois protection service available | boolean |  |
| ... authCode | Auth code required/needed/optional | text10 |  |
| ... minAuthLength | Minimal auth code length | int |  |
| ... maxAuthLength | Maximal auth code length | int |  |
| ... transferLock | TLD registry supports transfer lock | boolean |  |
| ... createAllowed | Domain creation is allowed | text10 |  |
| ... renewAllowed | Manual domain renewal is allowed | text10 |  |
| ... tradeAllowed | Domain trade is allowed | text10 |  |
| ... pushAllowed | Domain push is allowed | text10 |  |
| ... updateAllowed | Domain update is allowed | text10 |  |
| ... updateAuthcodeAllowed | Update domain auth code is allowed | text10 |  |
| ... transferAllowed | Domain transfer is allowed | text10 |  |
| ... transferOutAllowed | Domain transfer out ack is allowed | text10 |  |
| ... restoreAllowed | Domain restore is allowed | text10 |  |
| ... paperwork | Domain action requires additional paperwork | int |  |
| ... restorePossibility | Domain restore possibility | int |  |
| ... blacklist | Not allowed second level domain names | text |  |
| ... defaultRegistrationPeriod | Default Registration Period | text |  |
| ... defaultRenewalPeriod | Default Renewal Period | text |  |
| ... accreditation | Accreditation | text |  |
| ... tldAce | Unicode TLD | text |  |


#### 2.9.10. domain.getTldGroups {#domain.getTldGroups}

Get groups for TLD


##### 2.9.10.1. Input {#id232}

Table 2.120. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| tld | Top level domain | array_text64 | false |  |


##### 2.9.10.2. Output {#id233}

Table 2.121. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| groups | All TLD groups this domain is part of | array_identifier |  |


#### 2.9.11. domain.info {#domain.info}

Get domain details.


##### 2.9.11.1. Input {#id235}

Table 2.122. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | false |  |
| roId | Domain id (Repository Object Identifier) | int | false |  |
| wide | More detailed output | int | false | 1 |


##### 2.9.11.2. Output {#id236}

Table 2.123. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain | int |  |
| domain | Domain name | text64 |  |
| domain-ace | Domain ace name | text64 |  |
| period | Domain registration/renewal period | period | Yes |
| crDate | Time of domain creation | dateTime | Yes |
| exDate | Time of domain expiration | dateTime | Yes |
| upDate | Time of last domain update | dateTime | Yes |
| reDate | Time of domain renewal | dateTime | Yes |
| scDate | Time of scheduled execution | dateTime | Yes |
| transferLock | Domain is locked | boolean | Yes |
| status | Status of the domain | text64 | Yes |
| domainFlags | Domain status flags | array | Yes |
| authCode | Domain auth code | token255 | Yes |
| renewalMode | Domain renewal mode | renewalMode | Yes |
| transferMode | Domain transfer mode | transferMode | Yes |
| registrant | Domain owner contact handle id | int | Yes |
| admin | Domain administrative contact handle id | int | Yes |
| tech | Domain technical contact handle id | int | Yes |
| billing | Domain billing contact handle id | int | Yes |
| ns | List of nameserver | nsList | Yes |
| noDelegation | Domain delegation status | boolean | Yes |
| contact |  |  | Yes |
| ... registrant | Registrant contact handle details | contact |  |
| ... admin | Administrative contact handle details | contact |  |
| ... tech | Technical contact handle details | contact |  |
| ... billing | Billing contact handle details | contact |  |
| extData | Domain extra data | extData | Yes |
| verificationStatus | Contact Verification Status | text64 | Yes |
| registrantVerificationStatus | Contact Verification Status | text64 | Yes |


#### 2.9.12. domain.list {#domain.list}

List all available customers domains.


##### 2.9.12.1. Input {#id238}

Table 2.124. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Filter by domain name search string | array_text255 | false | * |
| roId | Filter by domain id (Repository Object Identifier) | array_int | false |  |
| status | Filter by domain status | array_text64 | false |  |
| registrant | Filter by registrant ids | array_int | false |  |
| admin | Filter by admin ids | array_int | false |  |
| tech | Filter by tech ids | array_int | false |  |
| billing | Filter by billing ids | array_int | false |  |
| renewalMode | Filter by renewal mode | renewalMode | false |  |
| transferLock | Filter by transfer lock | boolean | false |  |
| noDelegation | Filter by delegation status | boolean | false |  |
| tag | Filter by tag ids | array_int | false |  |
| wide | More detailed output | int | false | 1 |
| order | Sort order of result list | domainOrder | false | DOMAINASC |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |
| withPrivacy | Filter by Privacy option (1 is true) | int | false |  |


##### 2.9.12.2. Output {#id239}

Table 2.125. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of preregistrations | int |  |
| domain |  |  |  |
| ... roId | Id (Repository Object Identifier) of the domain | int |  |
| ... domain | Domain name | token255 |  |
| ... domain-ace | Domain ace name | token255 |  |
| ... period | Domain registration/renewal period | period | Yes |
| ... crDate | Time of domain creation | dateTime | Yes |
| ... exDate | Time of domain expiration | dateTime | Yes |
| ... upDate | Time of last domain update | dateTime | Yes |
| ... reDate | Time of domain renewal | dateTime | Yes |
| ... scDate | Time of scheduled execution | dateTime | Yes |
| ... transferLock | Domain is locked | boolean | Yes |
| ... status | Status of the domain | text64 | Yes |
| ... authCode | Domain auth code | token255 | Yes |
| ... renewalMode | Domain renewal mode | renewalMode | Yes |
| ... transferMode | Domain transfer mode | transferMode | Yes |
| ... registrant | Domain owner contact handle id | int | Yes |
| ... admin | Domain administrative contact handle id | int | Yes |
| ... tech | Domain technical contact handle id | int | Yes |
| ... billing | Domain billing contact handle id | int | Yes |
| ... ns | List of nameserver | nsList | Yes |
| ... noDelegation | Domain delegation status | boolean | Yes |
| ... contact |  |  | Yes |
| ... ... registrant | Registrant contact handle details | contact |  |
| ... ... admin | Administrative contact handle details | contact |  |
| ... ... tech | Technical contact handle details | contact |  |
| ... ... billing | Billing contact handle details | contact |  |
| ... extData | Domain extra data | extData | Yes |
| ... verificationStatus | Contact Verification Status | text64 | Yes |


#### 2.9.13. domain.log {#domain.log}

Log of changes to a domain.


##### 2.9.13.1. Input {#id241}

Table 2.126. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Filter result by domain name | token255 | false |  |
| status | Filter result by status | text64 | false |  |
| invoice | Filter result by invoice id | text64 | false |  |
| dateFrom | Log start date | timestamp | false |  |
| dateTo | Log end date | timestamp | false |  |
| priceMin | Minimum price of log entry | float | false | 0.0 |
| priceMax | Maximum price of log entry | float | false |  |
| order | ordering of the results | domainLogOrder | false | LOGTIMEDESC |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.9.13.2. Output {#id242}

Table 2.127. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of log entries | int |  |
| sum | Sum of amount | float |  |
| domain |  |  |  |
| ... customerId | ID of the customer | int |  |
| ... account | Account that performed the action | username |  |
| ... date | Log timestamp | dateTime |  |
| ... roId | Id (Repository Object Identifier) of the domain | int |  |
| ... domain | Domain name | token255 |  |
| ... status | Domain status | text64 |  |
| ... price | Incurred expenses of the domain action | float |  |
| ... invoice | Invoice id | text64 |  |
| ... remoteAddr | Ip address of executing client | ip |  |
| ... userText | Domain action description | text255 |  |
| ... logId | ID of the log entry | int |  |


#### 2.9.14. domain.priceChanges {#domain.priceChanges}

Get changes of domain prices.


##### 2.9.14.1. Input {#id244}

No parameters allowed


##### 2.9.14.2. Output {#id245}

Table 2.128. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| futureChanges |  |  |  |
| ... tld | Top level domain name | text64 |  |
| ... affectsCustomer | Customer has active domains for the TLD | boolean |  |
| ... currency | Currency of the price | text10 |  |
| ... changeDate | Date of the price's change | dateTime |  |
| ... createPrice | Domain registration price/year | float |  |
| ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... transferPrice | Domain transfer price/year | float |  |
| ... renewalPrice | Domain renewal price/year | float |  |
| ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... updatePrice | Domain update price | float |  |
| ... tradePrice | Domain trade price | float | Yes |
| ... trusteePrice | Domain trustee service price/year | float | Yes |
| ... monthlyTrusteePrice | Domain trustee service price/month | float | Yes |
| ... restorePrice | Domain restore price | float | Yes |
| ... chargedPrice |  |  | Yes |
| ... ... currency | Currency of the charged price | text10 |  |
| ... ... createPrice | Domain registration price/year | float |  |
| ... ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... ... transferPrice | Domain transfer price/year | float |  |
| ... ... renewalPrice | Domain renewal price/year | float |  |
| ... ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... ... updatePrice | Domain update price | float |  |
| ... ... tradePrice | Domain trade price | float | Yes |
| ... ... trusteePrice | Domain trustee service price/year | float | Yes |
| ... ... monthlyTrusteePrice | Domain trustee service price/month | float | Yes |
| ... ... restorePrice | Domain restore price | float | Yes |
| pastChanges |  |  |  |
| ... tld | Top level domain name | text64 |  |
| ... currency | Currency of the price | text10 |  |
| ... changeDate | Date of the price's change | dateTime |  |
| ... createPrice | Domain registration price/year | float |  |
| ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... transferPrice | Domain transfer price/year | float |  |
| ... renewalPrice | Domain renewal price/year | float |  |
| ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... updatePrice | Domain update price | float |  |
| ... tradePrice | Domain trade price | float | Yes |
| ... trusteePrice | Domain trustee service price/year | float | Yes |
| ... monthlyTrusteePrice | Domain trustee service price/month | float | Yes |
| ... restorePrice | Domain restore price | float | Yes |
| ... chargedPrice |  |  | Yes |
| ... ... currency | Currency of the charged price | text10 |  |
| ... ... createPrice | Domain registration price/year | float |  |
| ... ... monthlyCreatePrice | Domain registration price/month | float | Yes |
| ... ... transferPrice | Domain transfer price/year | float |  |
| ... ... renewalPrice | Domain renewal price/year | float |  |
| ... ... monthlyRenewalPrice | Domain renewal price/month | float | Yes |
| ... ... updatePrice | Domain update price | float |  |
| ... ... tradePrice | Domain trade price | float | Yes |
| ... ... trusteePrice | Domain trustee service price/year | float | Yes |
| ... ... monthlyTrusteePrice | Domain trustee service price/month | float | Yes |
| ... ... restorePrice | Domain restore price | float | Yes |


#### 2.9.15. domain.push {#domain.push}

Return domain to registry (if supported by TLD).


##### 2.9.15.1. Input {#id247}

Table 2.129. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| target | Target registrar (if supported) | text64 | false |  |
| scDate | Scheduled execution date | timestamp | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.15.2. Output {#id248}

No additional return parameters


#### 2.9.16. domain.removeClientHold {#domain.removeClientHold}

Remove client hold status for specified domain. Removing the clientHold status is only possible if it was set by you.


##### 2.9.16.1. Input {#id250}

Table 2.130. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.16.2. Output {#id251}

Table 2.131. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain | int |  |


#### 2.9.17. domain.renew {#domain.renew}

Manual domain renewal.


##### 2.9.17.1. Input {#id253}

Table 2.132. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| period | Domain renewal period | period | true |  |
| expiration | Date of current expiration | date | true |  |
| asynchron | Asynchron domain renew | boolean | false | false |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.17.2. Output {#id254}

Table 2.133. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain | int |  |
| price | Incurred expenses | float |  |
| currency | Currency related to price value | customercurrency |  |
| exDate | Time of domain expiration | dateTime |  |
| reDate | Time of next domain renewal | dateTime |  |


#### 2.9.18. domain.restore {#domain.restore}

Restores an expired/deleted domain (if supported).


##### 2.9.18.1. Input {#id256}

Table 2.134. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| renewalMode | Domain renewal mode | renewalMode | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.18.2. Output {#id257}

Table 2.135. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| domain | Domain name | text64 |  |
| domainAce | Domain ace name | text64 |  |


#### 2.9.19. domain.setClientHold {#domain.setClientHold}

Set client hold status for specified domain.


##### 2.9.19.1. Input {#id259}

Table 2.136. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.19.2. Output {#id260}

Table 2.137. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain | int |  |


#### 2.9.20. domain.stats {#domain.stats}

Get registered domain TLD statistics.


##### 2.9.20.1. Input {#id262}

No parameters allowed


##### 2.9.20.2. Output {#id263}

Table 2.138. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| tld |  |  | Yes |
| ... count | Total number of registered TLD | int |  |


#### 2.9.21. domain.trade {#domain.trade}

Change domain owner (if supported).


##### 2.9.21.1. Input {#id265}

Table 2.139. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| registrant | New domain owner contact handle id | int | true |  |
| admin | New administrative contact handle id | int | false |  |
| tech | New technical contact handle id | int | false |  |
| billing | New billing contact handle id | int | false |  |
| ns | List of nameserver | nsList | false |  |
| authCode | Authorization code (if supported) | token255 | false |  |
| whoisProvider | Whois provider | token0255 | false |  |
| whoisUrl | Whois url | token0255 | false |  |
| scDate | Scheduled execution date | timestamp | false |  |
| extData | Domain trade extra data | extData | false |  |
| asynchron | Asynchron domain trade | boolean | false | false |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.21.2. Output {#id266}

Table 2.140. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain | int |  |
| price | Incurred expenses | float |  |
| currency | Currency related to price value | customercurrency |  |
| notifications | Information about not required contact types | array | Yes |


#### 2.9.22. domain.transfer {#domain.transfer}

Transfer domain from another registrar or user to you.


##### 2.9.22.1. Input {#id268}

Table 2.141. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| registrant | Domain owner contact handle id | int | false |  |
| admin | Domain administrative contact handle id | int | false |  |
| tech | Domain technical contact handle id | int | false |  |
| billing | Domain billing contact handle id | int | false |  |
| ns | List of nameserver | nsList | false |  |
| nsTakeover | Keep existing nameservers | boolean | false | false |
| contactTakeover | Transfer also data related to the transfered gTLD | boolean | false | false |
| transferLock | Lock domain | boolean | false | 1 |
| authCode | Authorization code (if supported) | token255 | false |  |
| renewalMode | Domain renewal mode | renewalMode | false |  |
| whoisProvider | Whois provider | token0255 | false |  |
| whoisUrl | Whois url | token0255 | false |  |
| extData | Domain extra data | extData | false |  |
| scDate | Time of scheduled execution | timestamp | false |  |
| asynchron | Asynchron domain transfer | boolean | false | false |
| voucher | Voucher code | text64 | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.22.2. Output {#id269}

Table 2.142. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the domain | int |  |
| price | Incurred expenses | float |  |
| currency | Currency related to price value | customercurrency |  |
| notifications | Information about not required contact types | array | Yes |


#### 2.9.23. domain.transfercancel {#domain.transfercancel}

Cancel the transfer of a domain to another registrar or user.


##### 2.9.23.1. Input {#id271}

Table 2.143. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |


##### 2.9.23.2. Output {#id272}

No additional return parameters


#### 2.9.24. domain.transferOut {#domain.transferOut}

Allow or deny outgoing transfer requests.


##### 2.9.24.1. Input {#id274}

Table 2.144. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| answer | Acknowledge or deny the domain request | transferAnswer | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.24.2. Output {#id275}

No additional return parameters


#### 2.9.25. domain.update {#domain.update}

Update domain data.


##### 2.9.25.1. Input {#id277}

Table 2.145. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| registrant | New owner contact handle id | int | false |  |
| admin | New administrative contact handle id | int | false |  |
| tech | New technical contact handle id | int | false |  |
| billing | New billing contact handle id | int | false |  |
| ns | List of nameserver | nsList | false |  |
| transferLock | Lock domain | boolean | false |  |
| period | Domain registration/renewal period | period | false |  |
| authCode | Authorization code (if supported) | token255 | false |  |
| scDate | Time of scheduled execution | timestamp | false |  |
| renewalMode | Domain renewal mode | renewalMode | false |  |
| transferMode | Domain transfer mode | transferMode | false |  |
| whoisProvider | Whois provider | token0255 | false |  |
| whoisUrl | Whois url | token0255 | false |  |
| extData | Domain extra data | extData | false |  |
| asynchron | Asynchron domain update | boolean | false | false |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.9.25.2. Output {#id278}

Table 2.146. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| price | Incurred expenses | float |  |
| currency | Currency related to price value | customercurrency |  |
| notifications | Information about not required contact types | array | Yes |


#### 2.9.26. domain.whois {#domain.whois}

Get the whois information of a domain.


##### 2.9.26.1. Input {#id280}

Table 2.147. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |


##### 2.9.26.2. Output {#id281}

Table 2.148. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| whois | Whois output | text |  |


## 2.10. Dyndns {#id11241}

The dyndns object provides methods to manage DynDNS accounts.


#### 2.10.1. dyndns.changepassword {#dyndns.changepassword}

Change current password.


##### 2.10.1.1. Input {#id283}

Table 2.149. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| username | Customers username | text64 | true |  |
| password | Customers password | password | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.10.1.2. Output {#id284}

No additional return parameters


#### 2.10.2. dyndns.check {#dyndns.check}

Check DynDns Accounts


##### 2.10.2.1. Input {#id286}

Table 2.150. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| username | Username | username | false |  |
| hostname | A valid hostname | array_text255 | false |  |


##### 2.10.2.2. Output {#id287}

Table 2.151. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| username | Total number of preregistrations | text255 |  |
| available | Is hostname or username available | boolean |  |


#### 2.10.3. dyndns.create {#dyndns.create}

Create DynDns Account


##### 2.10.3.1. Input {#id289}

Table 2.152. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| username | Username | username | true |  |
| password | Password | text64 | true |  |
| hostname | A valid hostname | array_text255 | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.10.3.2. Output {#id290}

Table 2.153. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| accountId | ID of your account | int |  |
| recordId | ID of your DNS record | int |  |


#### 2.10.4. dyndns.delete {#dyndns.delete}

Delete a DynDns Account


##### 2.10.4.1. Input {#id292}

Table 2.154. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| accountId | Account-ID of a DynDNS Account | int | false |  |
| username | Username of a DynDNS Account | username | false |  |
| testing | Testing mode will perform no real action | boolean | false |  |


##### 2.10.4.2. Output {#id293}

No additional return parameters


#### 2.10.5. dyndns.info {#dyndns.info}

Info about a DynDns Account


##### 2.10.5.1. Input {#id295}

Table 2.155. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| accountId | Account-ID of a DynDNS Account | int | false |  |
| username | Username of a DynDNS Account | username | false |  |


##### 2.10.5.2. Output {#id296}

Table 2.156. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| accountId | ID of your account | int |  |
| created | Create date | dateTime |  |
| username | Your username | text255 |  |
| records | DNS records | array |  |


#### 2.10.6. dyndns.list {#dyndns.list}

List of your DynDns Accounts


##### 2.10.6.1. Input {#id298}

Table 2.157. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| username | Username | username | false |  |
| roid | roId of your DynDns Account | int | false |  |


##### 2.10.6.2. Output {#id299}

Table 2.158. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| accountId | ID of your account | int |  |
| created | Create date | dateTime |  |
| username | Your username | text255 |  |
| records | List of DNS records | array |  |


#### 2.10.7. dyndns.log {#dyndns.log}

Recent log-messages of a DynDns Account


##### 2.10.7.1. Input {#id301}

Table 2.159. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| accountId | Account-ID of a DynDNS Account | int | true |  |


##### 2.10.7.2. Output {#id302}

Table 2.160. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Number of logs | int |  |
| logs | Logs | array |  |


#### 2.10.8. dyndns.updateRecord {#dyndns.updateRecord}

Updates a DynDNS record.


##### 2.10.8.1. Input {#id304}

Table 2.161. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| ipAddress | A valid IPv4 address (A record) | ip | false |  |
| ipAddressV6 | A valid IPv6 address (AAAA record, optional) | ip | false |  |


##### 2.10.8.2. Output {#id305}

Table 2.162. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| ipv4_updated | Number of IPv4 updates | int |  |
| ipv6_updated | Number of IPv6 updates | int |  |


## 2.11. DyndnsSubscription {#id11728}

The dyndnssubscription object provides methods to manage the DynDNS-subscriptions.


#### 2.11.1. dyndnssubscription.cancel {#dyndnssubscription.cancel}

Cancels a DynDNS-subscription


##### 2.11.1.1. Input {#id307}

Table 2.163. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Customer's DynDNS-subscription id | int | true |  |


##### 2.11.1.2. Output {#id308}

No additional return parameters


#### 2.11.2. dyndnssubscription.create {#dyndnssubscription.create}

Order a DynDNS-subscription


##### 2.11.2.1. Input {#id310}

Table 2.164. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Id of the DynDNS-subscription plan | int | true |  |


##### 2.11.2.2. Output {#id311}

No additional return parameters


#### 2.11.3. dyndnssubscription.list {#dyndnssubscription.list}

Lists all active DynDNS-subscriptions of the customer


##### 2.11.3.1. Input {#id313}

Table 2.165. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.11.3.2. Output {#id314}

Table 2.166. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| subscription |  |  |  |
| ... id | Id of the DynDNS-subscription | int |  |
| ... name | Name of the DynDNS-subscription | text255 |  |
| ... accountsAmount | Amount of DynDNS-accounts | int |  |
| ... price | Price of the DynDNS-subscription for a 1 year term (without VAT) | float |  |
| ... currency | Currency of the price | text255 |  |
| ... createdAt | Date of the creation of the order | dateTime |  |
| ... lastPaymentAt | Date of the last payment for the DynDNS-subscription | dateTime |  |
| ... paidUntil | Expiration date of the DynDNS-subscription | dateTime |  |


#### 2.11.4. dyndnssubscription.listProducts {#dyndnssubscription.listProducts}

Lists all purchasable DynDNS-subscription plans with prices


##### 2.11.4.1. Input {#id316}

Table 2.167. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.11.4.2. Output {#id317}

Table 2.168. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| subscription |  |  |  |
| ... id | Id of the DynDNS-subscription plan | int |  |
| ... name | Name of the DynDNS-subscription plan | text255 |  |
| ... accountsAmount | Amount of DynDNS-accounts | int |  |
| ... price | Price of the DynDNS-subscription for a 1 year term (without VAT) | float |  |
| ... currency | Currency of the price | text255 |  |


## 2.12. Host {#id11976}

The host object provides methods to manage your glue records.


#### 2.12.1. host.check {#host.check}

Checks a hostname.


##### 2.12.1.1. Input {#id319}

Table 2.169. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostname | Name of host | hostname | true |  |


##### 2.12.1.2. Output {#id320}

Table 2.170. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| avail | Host available | boolean |  |


#### 2.12.2. host.create {#host.create}

Creates a new host.


##### 2.12.2.1. Input {#id322}

Table 2.171. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostname | Name of host | hostname | true |  |
| ip | Ip address(es) | array_ip | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.12.2.2. Output {#id323}

Table 2.172. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the hostname | int |  |


#### 2.12.3. host.delete {#host.delete}

Deletes a host.


##### 2.12.3.1. Input {#id325}

Table 2.173. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostname | Name of host | hostname | false |  |
| roId | Id (Repository Object Identifier) of the hostname | int | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.12.3.2. Output {#id326}

No additional return parameters


#### 2.12.4. host.info {#host.info}

Get host details.


##### 2.12.4.1. Input {#id328}

Table 2.174. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostname | Name of host | hostname | false |  |
| roId | Id (Repository Object Identifier) of the hostname | int | false |  |


##### 2.12.4.2. Output {#id329}

Table 2.175. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the hostname | int |  |
| hostname | Name of host | int |  |
| status | Status of the hostname | text64 |  |
| ip | List of ip address | ipList |  |


#### 2.12.5. host.list {#host.list}

List of hosts.


##### 2.12.5.1. Input {#id331}

Table 2.176. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| search | Filter by hostname search string | text64 | false | * |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.12.5.2. Output {#id332}

Table 2.177. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of hostnames | int |  |
| host |  |  |  |
| ... roId | Id (Repository Object Identifier) of the hostname | int |  |
| ... hostname | Name of host | token255 |  |
| ... ip | List of ip address | ipList |  |
| ... status | Status of the hostname | text64 |  |


#### 2.12.6. host.update {#host.update}

Updates a hostname.


##### 2.12.6.1. Input {#id334}

Table 2.178. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostname | Name of host | hostname | false |  |
| roId | Id (Repository Object Identifier) of the hostname | int | false |  |
| ip | Ip address(es) | array_ip | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.12.6.2. Output {#id335}

No additional return parameters


## 2.13. Hosting {#id12347}

The hosting object provides methods to manage the hosting package.


#### 2.13.1. hosting.cancel {#hosting.cancel}

Cancel a hosting package.


##### 2.13.1.1. Input {#id337}

Table 2.179. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostingId | Hosting package id | token255 | true |  |


##### 2.13.1.2. Output {#id338}

No additional return parameters


#### 2.13.2. hosting.controlPanel {#hosting.controlPanel}

Returns the URL to the hosting control panel.


##### 2.13.2.1. Input {#id340}

Table 2.180. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Hosting package ID | int | true |  |


##### 2.13.2.2. Output {#id341}

Table 2.181. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| url | URL to the hosting control panel | token255 |  |


#### 2.13.3. hosting.create {#hosting.create}

Order a hosting package.


##### 2.13.3.1. Input {#id343}

Table 2.182. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostingId | Hosting package id | token255 | false |  |
| productId | Hosting package id | token255 | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.13.3.2. Output {#id344}

No additional return parameters


#### 2.13.4. hosting.getPrices {#hosting.getPrices}

Get hosting packages prices.


##### 2.13.4.1. Input {#id346}

No parameters allowed


##### 2.13.4.2. Output {#id347}

Table 2.183. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| hosting |  |  |  |
| ... id | Id of the hosting package | int | Yes |
| ... product | Name of the hosting package product | text64 | Yes |


#### 2.13.5. hosting.issuspended {#hosting.issuspended}

Check whether a hosting plan is currently suspended


##### 2.13.5.1. Input {#id349}

Table 2.184. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostingId | ID of the hosting product | int | true |  |


##### 2.13.5.2. Output {#id350}

Table 2.185. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| isSuspended | Indicates whether the specified hosting plan is currently suspended | boolean |  |
| endDate | Cancellation time of the specified hosting plan | dateTime |  |
| paidUntil | Paid until information for the hosting product | dateTime |  |


#### 2.13.6. hosting.list {#hosting.list}

Lists customers hosting packages.


##### 2.13.6.1. Input {#id352}

No parameters allowed


##### 2.13.6.2. Output {#id353}

Table 2.186. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of hosting packages | int |  |
| hosting |  |  |  |
| ... id | Id of the hosting package | int | Yes |
| ... product | Name of the hosting package product | text64 | Yes |
| ... cancellationDate | Date of cancellation of action | dateTime | Yes |
| ... autorechnung_id | Automatic readout id | int | Yes |
| ... product_id | Hosting package id | int | Yes |
| ... serverNo | Customers company server number | int | Yes |
| ... serverAccount | The Customers account number | int | Yes |
| ... serverIP | Ip address of the server number | ip | Yes |
| ... serverMail | Email address of the server number | email | Yes |
| ... producttype | Type of the hosting package product | text64 | Yes |
| ... suspended | Whether the hosting plan is currently suspended | boolean | Yes |
| ... pma |  | int | Yes |
| ... webMail |  | email | Yes |
| ... ftp |  | ip | Yes |
| ... smtp |  | ip | Yes |
| ... smtpIn0 |  | ip | Yes |
| ... smtpIn1 |  | ip | Yes |
| ... externIp | IP of the hosting | ip | Yes |
| ... externIpPort | IP of the hosting PORT | ip | Yes |
| ... ipv6 | Internet Protocol version 6 (IPv6) | ip | Yes |
| ... nextCancellationDate | Next date of service cancellation | dateTime | Yes |
| ... cancellationRefund | Refund amount | float | Yes |
| ... executionPeriod | Current billing period in months | int | Yes |


#### 2.13.7. hosting.reinstate {#hosting.reinstate}

Reinstate a hosting package.


##### 2.13.7.1. Input {#id355}

Table 2.187. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostingId | Hosting package id | token255 | true |  |


##### 2.13.7.2. Output {#id356}

No additional return parameters


#### 2.13.8. hosting.unsuspend {#hosting.unsuspend}

Unsuspend a suspend hosting product of a customer


##### 2.13.8.1. Input {#id358}

Table 2.188. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostingId | ID of the hosting product | int | true |  |
| testing | Execute order in testing mode (no order will be submitted) | boolean | false | false |


##### 2.13.8.2. Output {#id359}

Table 2.189. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| netPrice | Net price for restoring and extending hosting plan | float |  |
| grossPrice | Gross price for restoring and extending hosting plan | float |  |
| currency | Currency of paid price | text255 |  |
| quantity | Quantity | int |  |
| unit | Unit (months or years) | text255 |  |
| message | System message | text255 |  |


#### 2.13.9. hosting.updatePeriod {#hosting.updatePeriod}

Update the billing period of a hosting package.


##### 2.13.9.1. Input {#id361}

Table 2.190. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| hostingId | ID of the hosting product | int | true |  |
| period | New billing period (1, 3, 6, 12, 18, 24 months, but not less than default for the product) | int | true |  |


##### 2.13.9.2. Output {#id362}

No additional return parameters


## 2.14. Message {#id12887}

The message object provides methods to query the message queue.


#### 2.14.1. message.ack {#message.ack}

Acknowledge and remove a message from the notification queue.


##### 2.14.1.1. Input {#id364}

Table 2.191. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Message ID which can be obtained via the poll method | int | true |  |


##### 2.14.1.2. Output {#id365}

No additional return parameters


#### 2.14.2. message.poll {#message.poll}

Get the first message from your notification queue.


##### 2.14.2.1. Input {#id367}

No parameters allowed


##### 2.14.2.2. Output {#id368}

Table 2.192. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Number of unread messages | int |  |
| msg |  |  |  |
| ... type | Type of the message | message_type |  |
| ... id | ID of the message | int |  |
| ... date | Time of the message creation date | dateTime |  |
| ... object | Name of object, for example a domain name | token255 |  |
| ... status | Status of the object | message_status |  |
| ... statusDetails | Extra information on for example why an action did fail | text | Yes |


## 2.15. Nameserver {#id12987}

The nameserver object provides methods to manage the nameserver domains and their records.


#### 2.15.1. nameserver.check {#nameserver.check}

Checks if the given nameservers are responding accordingly.


##### 2.15.1.1. Input {#id370}

Table 2.193. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| ns | List of nameserver | nsList | true |  |


##### 2.15.1.2. Output {#id371}

Table 2.194. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| details |  |  |  |
| ... ns | Given nameserver | token255 |  |
| ... messageList | Check messages | array_text255 |  |
| ... status | Stauts of the nameserver check | text64 |  |


#### 2.15.2. nameserver.clone {#nameserver.clone}

Clones cource domain DNS to target DNS.


##### 2.15.2.1. Input {#id373}

Table 2.195. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| sourceDomain | Source domain name | token255 | true |  |
| targetDomain | Target domain name | token255 | true |  |


##### 2.15.2.2. Output {#id374}

Table 2.196. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | ID of new zone | int |  |


#### 2.15.3. nameserver.create {#nameserver.create}

Creates a domain in the nameserver.


##### 2.15.3.1. Input {#id376}

Table 2.197. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |
| type | Type of nameserver entry | nsType | true |  |
| ns | List of nameserver | nsList | false |  |
| masterIp | Master IP address | ip | false |  |
| web | Web nameserver entry | token0255 | false |  |
| mail | Mail nameserver entry | text0255 | false |  |
| soaEmail | Email address for SOA record | email | false |  |
| urlRedirectType | Type of the url redirection | urlRedirectType | false |  |
| urlRedirectTitle | Title of the frame redirection | token0255 | false |  |
| urlRedirectDescription | Description of the frame redirection | token0255 | false |  |
| urlRedirectFavIcon | FavIcon of the frame redirection | token0255 | false |  |
| urlRedirectKeywords | Keywords of the frame redirection | token0255 | false |  |
| testing | Execute command in testing mode | boolean | false | false |
| ignoreExisting | Ignore existing | boolean | false | false |


##### 2.15.3.2. Output {#id377}

Table 2.198. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Created DNS domain id | int |  |


#### 2.15.4. nameserver.createRecord {#nameserver.createRecord}

Creates a new nameserver record.


##### 2.15.4.1. Input {#id379}

Table 2.199. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | false |  |
| roId | DNS domain id | int | false |  |
| type | Type of the nameserver record | recordType | true |  |
| content | Content of the nameserver record | text | true |  |
| name | Name of the nameserver record | token0255 | false |  |
| ttl | TTL (time to live) of the nameserver record | int | false | 3600 |
| prio | Priority of the nameserver record | int | false | 0 |
| urlRedirectType | Type of the url redirection | urlRedirectType | false |  |
| urlRedirectTitle | Title of the frame redirection | token0255 | false |  |
| urlRedirectDescription | Description of the frame redirection | token0255 | false |  |
| urlRedirectFavIcon | FavIcon of the frame redirection | token0255 | false |  |
| urlRedirectKeywords | Keywords of the frame redirection | token0255 | false |  |
| urlAppend | Append the path for redirection | boolean | false | false |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.15.4.2. Output {#id380}

Table 2.200. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | Created record id | int |  |


#### 2.15.5. nameserver.delete {#nameserver.delete}

Deletes a nameserver domain


##### 2.15.5.1. Input {#id382}

Table 2.201. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | false |  |
| roId | Id (Repository Object Identifier) of the DNS domain | int | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.15.5.2. Output {#id383}

No additional return parameters


#### 2.15.6. nameserver.deleteRecord {#nameserver.deleteRecord}

Deletes a nameserver record.


##### 2.15.6.1. Input {#id385}

Table 2.202. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Id of the nameserver record | int | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.15.6.2. Output {#id386}

No additional return parameters


#### 2.15.7. nameserver.export {#nameserver.export}

Creates a nameserver.export TXT Datei


##### 2.15.7.1. Input {#id388}

Table 2.203. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | true |  |


##### 2.15.7.2. Output {#id389}

Table 2.204. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| export | export result | text |  |


#### 2.15.8. nameserver.exportlist {#nameserver.exportlist}

Nameserver List export as file


##### 2.15.8.1. Input {#id391}

Table 2.205. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| format | Format of the requested document | documentformat | false | raw |
| domain | Search by domain name | token0255 | false | * |
| wide | More detailed output | int | false | 1 |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.15.8.2. Output {#id392}

Table 2.206. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Log timestamp | int |  |
| domains | Domains Log list | array |  |


#### 2.15.9. nameserver.exportrecords {#nameserver.exportrecords}

Export nameserver records as file


##### 2.15.9.1. Input {#id394}

Table 2.207. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| format | Request output of given format. Defaults to csv. | documentformat | false | csv |
| name | Filter results by given record name. You can, for example, set domain names here. Wildcard * is allowed. | token255 | false |  |
| page | If limit is set, show entries from {page - 1}*{limit}. Defaults to 1. | int | false | 1 |
| limit | Only return {limit} entries at once. Set to 0 for no limit. Defaults to 0. | int | false | 0 |


##### 2.15.9.2. Output {#id395}

Table 2.208. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Amount of entries returned. | int |  |
| data | Result file data in cleartext or base64 coded, depending on the selected format. CSV returns as clear text. Other formats are currently not supported. | base64 |  |


#### 2.15.10. nameserver.info {#nameserver.info}

Get nameserver record details. The request requires either the domain or roid.


##### 2.15.10.1. Input {#id397}

Table 2.209. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Search by domain name. The request requires either the domain or roid. | token255 | false |  |
| roId | Id (Repository Object Identifier) of the DNS domain. The request requires either the domain or roid. | int | false |  |
| recordId | Search by record id | int | false |  |
| type | Search by record type | text64 | false |  |
| name | Search by record name | text0255 | false |  |
| content | Search by record content | text1024 | false |  |
| ttl | Search by record ttl | int | false |  |
| prio | Search by record priority | int | false |  |


##### 2.15.10.2. Output {#id398}

Table 2.210. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| roId | Id (Repository Object Identifier) of the DNS domain | int | Yes |
| domain | Domain name | token255 | Yes |
| type | Type of nameserver domain | nsType | Yes |
| masterIp | Master IP address | ip | Yes |
| lastZoneCheck | Time of last zone check | dateTime | Yes |
| slaveDns |  |  | Yes |
| ... name | Hostname of the nameserver | hostname |  |
| ... ip | Ip address of the nameserver | ip |  |
| SOAserial | SOA-RR serial | text064 | Yes |
| count | Total number of domain records | int | Yes |
| record |  |  | Yes |
| ... id | Id of the nameserver record | int |  |
| ... name | Name of the nameserver record | token255 |  |
| ... type | Type of the nameserver record | recordType |  |
| ... content | Content of the nameserver record | text1024 |  |
| ... ttl | TTL (time to live) of the nameserver record | int |  |
| ... prio | Priority of the nameserver record | int |  |
| ... urlRedirectType | Type of the url redirection | urlRedirectType | Yes |
| ... urlRedirectTitle | Title of the frame redirection | token255 | Yes |
| ... urlRedirectDescription | Description of the frame redirection | token255 | Yes |
| ... urlRedirectKeywords | Keywords of the frame redirection | token255 | Yes |
| ... urlRedirectFavIcon | FavIcon of the frame redirection | token255 | Yes |
| ... urlAppend | Append the path to redirection | boolean | Yes |


#### 2.15.11. nameserver.list {#nameserver.list}

List all nameserver domains.


##### 2.15.11.1. Input {#id400}

Table 2.211. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Search by domain name | token0255 | false | * |
| wide | More detailed output | int | false | 1 |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.15.11.2. Output {#id401}

Table 2.212. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of nameserver domains | int |  |
| domains |  |  |  |
| ... roId | Id (Repository Object Identifier) of the DNS domain | int |  |
| ... domain | Domain name | array_domain |  |
| ... type | Type of nameserver domain | nsType |  |
| ... masterIp | Master ip address | ip | Yes |
| ... mail | Mail nameserver entry | ip_url | Yes |
| ... web | Web nameserver entry | ip_url | Yes |
| ... url | Web forwarding url | ip_url | Yes |
| ... urlType | The redirect type of the forwarding url (only if `url` is set) | urlRedirectType | Yes |
| ... ipv4 | Web IPv4 address | ip | Yes |
| ... ipv6 | Web IPv6 address | ip | Yes |


#### 2.15.12. nameserver.update {#nameserver.update}

Updates a nameserver domain.


##### 2.15.12.1. Input {#id403}

Table 2.213. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name | token255 | false |  |
| roId | Id (Repository Object Identifier) of the DNS domain | int | false |  |
| type | Type of nameserver entry | nsType | false |  |
| masterIp | Master ip address | ip | false |  |
| ns | List of nameserver | nsList | false |  |
| web | Web nameserver entry | token255 | false |  |
| mail | Mail nameserver entry | token255 | false |  |
| urlRedirectType | Type of the url redirection | urlRedirectType | false |  |
| urlRedirectTitle | Title of the frame redirection | token0255 | false |  |
| urlRedirectDescription | Description of the frame redirection | token0255 | false |  |
| urlRedirectFavIcon | FavIcon of the frame redirection | token0255 | false |  |
| urlRedirectKeywords | Keywords of the frame redirection | token0255 | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.15.12.2. Output {#id404}

No additional return parameters


#### 2.15.13. nameserver.updateRecord {#nameserver.updateRecord}

Updates a nameserver record.


##### 2.15.13.1. Input {#id406}

Table 2.214. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Id of the record | array_int | true |  |
| name | Name of the nameserver record | token0255 | false |  |
| type | Type of the nameserver record | recordType | false |  |
| content | Content of the nameserver record | text | false |  |
| prio | Priority of the nameserver record | int | false |  |
| ttl | TTL (time to live) of the nameserver record | int | false |  |
| urlRedirectType | Type of the url redirection | urlRedirectType | false |  |
| urlRedirectTitle | Title of the frame redirection | token0255 | false |  |
| urlRedirectDescription | Description of the frame redirection | token0255 | false |  |
| urlRedirectFavIcon | FavIcon of the frame redirection | token0255 | false |  |
| urlRedirectKeywords | Keywords of the frame redirection | token0255 | false |  |
| urlAppend | Append the path for redirection | boolean | false | false |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.15.13.2. Output {#id407}

No additional return parameters


## 2.16. NameserverSet {#id14267}

The nameserverset object provides methods to manage (create, update, list etc.) your nameserver sets.


#### 2.16.1. nameserverset.create {#nameserverset.create}

Creates a new nameserver set.


##### 2.16.1.1. Input {#id409}

Table 2.215. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| type | Type of the nameserver set | nsSetType | true |  |
| name | Name of the nameserver set | token0255 | false |  |
| ns | List of nameservers | nsList | true |  |
| hostmaster | Email address of the hostmaster | email | false |  |
| visible | Should the nameserver set be visible | boolean | false | true |
| prio | Priority of the nameserver set | int | false | 0 |
| web | Web nameserver entry | ip_url | false |  |
| mail | Mail nameserver entry | text0255 | false |  |
| masterIp | Master IP address | ip | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.16.1.2. Output {#id410}

Table 2.216. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the created nameserver set | int |  |


#### 2.16.2. nameserverset.delete {#nameserverset.delete}

Deletes a nameserver set.


##### 2.16.2.1. Input {#id412}

Table 2.217. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | ID of the nameserver set | int | true |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.16.2.2. Output {#id413}

No additional return parameters


#### 2.16.3. nameserverset.info {#nameserverset.info}

Get details of the nameserver set.


##### 2.16.3.1. Input {#id415}

Table 2.218. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | ID of the nameserver set | int | true |  |


##### 2.16.3.2. Output {#id416}

Table 2.219. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the nameserver set | token255 |  |
| roId | ID of the nameserver set | int |  |
| name | Name of the nameserver set | token0255 | Yes |
| readOnly | Is the set read only | boolean |  |
| type | Type of the nameserver set | nsSetType |  |
| ns | List of nameservers | nsList |  |
| hostmaster | Email address of the hostmaster | email |  |
| visible | Is the set visible | boolean |  |
| prio | Priority of the nameserver set | int | Yes |
| web | Web nameserver entry | ip_url | Yes |
| mail | Mail nameserver entry | text0255 | Yes |
| masterIp | Master IP address | ip | Yes |


#### 2.16.4. nameserverset.list {#nameserverset.list}

List all nameserver sets.


##### 2.16.4.1. Input {#id418}

Table 2.220. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| readOnly | List only readonly contact handle | boolean | false |  |
| wide | More detailed output, adds all existing optional return parameters | boolean | false | false |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.16.4.2. Output {#id419}

Table 2.221. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of nameserver sets | int |  |
| nsset |  |  |  |
| ... id | ID of the nameserver set | int |  |
| ... roId | ID of the nameserver set | int |  |
| ... name | Name of the nameserver set | token0255 | Yes |
| ... ns | List of nameservers | nsList |  |
| ... readOnly | Is the set read only | boolean | Yes |
| ... type | Type of nameserver set | nsSetType | Yes |
| ... hostmaster | Email address of the hostmaster | email | Yes |
| ... visible | Is the set visible | boolean | Yes |
| ... prio | Priority of the nameserver set | int | Yes |
| ... web | Web nameserver entry | ip_url | Yes |
| ... mail | Mail nameserver entry | text0255 | Yes |
| ... masterIp | Master IP address | ip | Yes |


#### 2.16.5. nameserverset.update {#nameserverset.update}

Udates an existing nameserver set.


##### 2.16.5.1. Input {#id421}

Table 2.222. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | ID of the nameserver set | int | true |  |
| type | Type of the nameserver set | nsSetType | false |  |
| name | Name of the nameserver set | token0255 | false |  |
| ns | List of nameservers | nsList | false |  |
| hostmaster | Email address of the hostmaster | email | false |  |
| visible | Should the nameserver set be visible | boolean | false |  |
| prio | Priority of the nameserver set | int | false |  |
| web | Web nameserver entry | ip_url | false |  |
| mail | Mail nameserver entry | text0255 | false |  |
| masterIp | Master IP address | ip | false |  |
| testing | Execute command in testing mode | boolean | false | false |


##### 2.16.5.2. Output {#id422}

No additional return parameters


## 2.17. News {#id14777}

The news object provides methods to list news feeds.


#### 2.17.1. news.list {#news.list}

Get the latest news


##### 2.17.1.1. Input {#id424}

Table 2.223. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | Get news with a specified ID | int | false |  |
| year | Get only news from the specified year | int | false |  |
| language | Get only news in the specified language | language | false |  |
| kdnr | Get only news for a specific customer and filter out all other INWX subsidiaries | int | false |  |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |


##### 2.17.1.2. Output {#id425}

Table 2.224. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of news | int |  |
| years | List of years that have at least one news item | array_int |  |
| news |  |  |  |
| ... id | Id of the news item | int |  |
| ... date | Date of the news | dateTime |  |
| ... title | Title of the news | token255 |  |
| ... textShort | Short text of the news | text |  |
| ... textLong | Long text of the news | text |  |
| provid | Current customers providerID (INWX subsidiary ID) | int |  |
| providerIds | When searching by a specific news ID, this array defines for which providerIDs the news exists | array_int |  |


## 2.18. Nichandle {#id14917}

The nichandle object provides methods to list customer's NIC handles


#### 2.18.1. nichandle.list {#nichandle.list}

Lists customers NIC handles


##### 2.18.1.1. Input {#id427}

Table 2.225. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| page | Page number for paging | int | false | 1 |
| pagelimit | Max number of results per page | int | false | 20 |
| handles | List of contact id to check NIC handles status | array_int | false |  |


##### 2.18.1.2. Output {#id428}

Table 2.226. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| NIC handles |  |  |  |
| ... status | Registry status of NIC handle | text064 |  |
| ... contactId | Contact id (handle id) | int |  |
| ... type | Type of NIC handle | text064 |  |
| ... domain | Example of an associated domain | text255 |  |
| ... details | Status details for FAILED | text0 | Yes |
| count | Amount of NIC handles of the customer | int |  |


## 2.19. Pdf {#id15018}

The pdf object provides methods to get required (pdf) documents.


#### 2.19.1. pdf.document {#pdf.document}

Get a generic PDF document.


##### 2.19.1.1. Input {#id430}

Table 2.227. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| category | Category of the PDF document to get. Possible values currently: dsgvo | text64 | true |  |
| name | Name of the PDF document. Possible values currently: contract, contract_empty, appendix1, appendix2 | text64 | true |  |


##### 2.19.1.2. Output {#id431}

Table 2.228. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| pdf | Base64 encoded PDF document | base64 |  |


#### 2.19.2. pdf.get {#pdf.get}

Get a required PDF document for a specific domain.


##### 2.19.2.1. Input {#id433}

Table 2.229. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Domain name for which to get the PDF | text64 | true |  |


##### 2.19.2.2. Output {#id434}

Table 2.230. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| pdf | Base64 encoded PDF document | base64 |  |


## 2.20. Tag {#id15125}

The tag object provides methods to manage (create, update, list etc.) your tags.


#### 2.20.1. tag.create {#tag.create}

Creates a new tag or returns the ID, if it already exists


##### 2.20.1.1. Input {#id436}

Table 2.231. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| name | Name of the new tag | token255 | true |  |


##### 2.20.1.2. Output {#id437}

Table 2.232. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the created tag | int |  |
| name | Name of your tag after removing invalid characters | token255 |  |


#### 2.20.2. tag.delete {#tag.delete}

Untags all objects and deletes a tag


##### 2.20.2.1. Input {#id439}

Table 2.233. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | ID of the tag to delete | int | true |  |


##### 2.20.2.2. Output {#id440}

No additional return parameters


#### 2.20.3. tag.info {#tag.info}

Get tag info.


##### 2.20.3.1. Input {#id442}

Table 2.234. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | ID of the tag | int | true |  |


##### 2.20.3.2. Output {#id443}

Table 2.235. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the tag | int |  |
| name | Name of the tag | token255 |  |


#### 2.20.4. tag.list {#tag.list}

List all existing tags by searching via tag IDs or tagged domains.


##### 2.20.4.1. Input {#id445}

Table 2.236. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| domain | Filter by single or multiple domain names | array_domain | false |  |
| id | Filter by single or multiple tag IDs | array_int | false |  |


##### 2.20.4.2. Output {#id446}

Table 2.237. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| count | Total number of tags | int |  |
| tag |  |  |  |
| ... id | ID of the tag | int |  |
| ... name | Name of the tag | token255 |  |
| ... count |  |  | Yes |
| ... ... domain | Number of domains using this tag, only shown when searched with tag ID | int |  |


#### 2.20.5. tag.update {#tag.update}

Update the name of an existing tag or add/remove said tag to domains. Updating tag name has priority and ignores add/rem.


##### 2.20.5.1. Input {#id448}

Table 2.238. Parameters


| Parameter | Description | Type | Required | Default |
| --- | --- | --- | --- | --- |
| id | ID of the tag | int | true |  |
| name | New tag name | token255 | false |  |
| add | Tag to object, can be a list of IDs | tagUpdateAdd | false |  |
| rem | Untag from object, can be a list of IDs | tagUpdateRem | false |  |


##### 2.20.5.2. Output {#id449}

Table 2.239. Parameters


| Parameter | Description | Type | Optional |
| --- | --- | --- | --- |
| id | ID of the tag, only present when name was updated | int | Yes |
| name | Name of the tag, only present when name was updated | token255 | Yes |


## Chapter 3: Datatypes


## Chapter 3. Datatypes {#types}


#### 3.1. _true {#type._true}

Only a boolean true value is allowed


#### 3.2. addressTitle {#type.addresstitle}

Address title

- MISS
- MISTER
- COMPANY


#### 3.3. applicationOrder {#type.applicationorder}

Sort order of the application list

- DOMAINASC
- DOMAINDESC
- PRICEASC
- PRICEDESC
- REGISTRANTASC
- REGISTRANTDESC


#### 3.4. array {#type.array}

One or a list of array(s)


#### 3.5. array_domain {#type.array_domain}

One or a list of valid domain name(s)


#### 3.6. array_float {#type.array_float}

A list or one float value


#### 3.7. array_identifier {#type.array_identifier}

One or a list of valid identifier name(s). Only alphanumeric characters and underscore is allowed. It must not start with a number.


#### 3.8. array_int {#type.array_int}

A list or one integer value


#### 3.9. array_ip {#type.array_ip}

A list or one valid ip (v4 or v6) address


#### 3.10. array_text {#type.array_text}

A list or one string value with a minimum lenght of 2 and maximum length of 30 characters


#### 3.11. array_text255 {#type.array_text255}

A list or one string value with a maximum length of 255 characters


#### 3.12. array_text64 {#type.array_text64}

A list or one string value with a maximum length of 64 characters


#### 3.13. auDomainIdType {#type.audomainidtype}

.AU domain ID type

- ABN : Australian Business Number
- ACN : Australian Company Number
- ARBN : Australian Registered Body Number
- OTHER : Other


#### 3.14. auDomainRelation {#type.audomainrelation}

.AU domain relation

- 1 : 2LD Domain name is an exact match, acronym or abbreviation of the registrant’s company or trading name, organization or association name or trademark.
- 2 : 2LD Domain Name is closely and substantially connected to the registrant.


#### 3.15. auDomainRelationType {#type.audomainrelationtype}

.AU domain relation type

- Company : Company
- Registered Business : Registered Business
- Sole Trader : Sole Trader
- Partnership : Partnership
- Trademark Owner : Trademark Owner
- Pending TM Owner : Pending TM Owner
- Incorporated Association : Incorporated Association
- Club : Club
- Non-profit Organisation : Non-profit Organisation
- Charity : Charity
- Trade Union : Trade Union
- Industry Body : Industry Body
- Commercial Statutory Body : Commercial Statutory Body
- Political Party : Political Party
- Citizen/Resident : Citizen/Resident
- Trust : Trust


#### 3.16. auEligibilityIdType {#type.aueligibilityidtype}

.AU eligibility ID type

- ACN : ACN
- ABN : ABN
- VIC BN : VIC BN
- NSW BN : NSW BN
- SA BN : SA BN
- NT BN : NT BN
- WA BN : WA BN
- TAS BN : TAS BN
- ACT BN : ACT BN
- QLD BN : QLD BN
- TM : TM
- OTHER : OTHER


#### 3.17. base64 {#type.base64}

Base64 encoded file data


#### 3.18. bgAppNumber {#type.bgappnumber}

.BG application number


#### 3.19. boolean {#type.boolean}

A boolean value

- 0
- 1


#### 3.20. boolean_3 {#type.boolean_3}

A boolean 3 value

- yes
- no
- null


#### 3.21. brCnpj {#type.brcnpj}

Brasilian CNPJ


#### 3.22. brCpf {#type.brcpf}

Brasilian CPF


#### 3.23. caLegalType {#type.calegaltype}

.CA Legal Type

- CCO : Corporation (Canada or Canadian province or territory)
- CCT : Canadian citizen
- RES : Permanent Resident of Canada
- GOV : Government or government entity in Canada
- EDU : Canadian Educational Institution
- ASS : Canadian Unincorporated Association
- HOP : Canadian Hospital
- PRT : Partnership Registered in Canada
- TDM : Trade-mark registered in Canada (by a non-Canadian owner)
- TRD : Canadian Trade Union
- PLT : Canadian Political Party
- LAM : Canadian Library, Archive or Museum
- TRS : Trust established in Canada
- ABO : Aboriginal Peoples (individuals or groups) indigenous to Canada
- INB : Indian Band recognized by the Indian Act of Canada
- LGR : Legal Representative of a Canadian Citizen or Permanent Resident
- OMK : Official mark registered in Canada
- MAJ : Her Majesty the Queen


#### 3.24. contact {#type.contact}

Contact data

- roId (int)
- id (text64)
- type (contactType)
- name (text64)
- org (text64)
- street (text64)
- city (text64)
- pc (text10)
- sp (text064)
- cc (country)
- voice (phone)
- fax (phone)
- email (email)
- remarks (text255)


#### 3.25. contactOrder {#type.contactorder}

Sort order of the contact list result

- IDDESC
- IDASC
- NAMEDESC
- NAMEASC


#### 3.26. contactType {#type.contacttype}

Allowed types of contact

- ORG
- PERSON
- ROLE


#### 3.27. country {#type.country}

Country in 2 letter country codes (ISO-3166-1)

- CW : CURAÇAO
- SX : SINT MAARTEN
- BQ : BONAIRE, SINT EUSTATIUS AND SABA
- XK : KOSOVO
- AF : AFGHANISTAN
- AX : ALAND ISLANDS
- AL : ALBANIA
- DZ : ALGERIA
- AS : AMERICAN SAMOA
- AD : ANDORRA
- AO : ANGOLA
- AI : ANGUILLA
- AQ : ANTARCTICA
- AG : ANTIGUA AND BARBUDA
- AR : ARGENTINA
- AM : ARMENIA
- AW : ARUBA
- AC : ASCENSION ISLAND
- AU : AUSTRALIA
- AT : AUSTRIA
- AZ : AZERBAIJAN
- BS : BAHAMAS
- BH : BAHRAIN
- BD : BANGLADESH
- BB : BARBADOS
- BY : BELARUS
- BE : BELGIUM
- BZ : BELIZE
- BJ : BENIN
- BM : BERMUDA
- BT : BHUTAN
- BO : BOLIVIA
- BA : BOSNIA AND HERZEGOVINA
- BW : BOTSWANA
- BV : BOUVET ISLAND
- BR : BRAZIL
- IO : BRITISH INDIAN OCEAN TERRITORY
- BN : BRUNEI DARUSSALAM
- BG : BULGARIA
- BF : BURKINA FASO
- BI : BURUNDI
- KH : CAMBODIA
- CM : CAMEROON
- CA : CANADA
- CV : CAPE VERDE
- KY : CAYMAN ISLANDS
- CF : CENTRAL AFRICAN REPUBLIC
- TD : CHAD
- CL : CHILE
- CN : CHINA
- CX : CHRISTMAS ISLAND
- CC : COCOS (KEELING) ISLANDS
- CO : COLOMBIA
- KM : COMOROS
- CG : CONGO
- CD : CONGO, THE DEMOCRATIC REPUBLIC OF THE
- CK : COOK ISLANDS
- CR : COSTA RICA
- CI : COTE D IVOIRE
- HR : CROATIA
- CU : CUBA
- CY : CYPRUS
- CZ : CZECH REPUBLIC
- DK : DENMARK
- DJ : DJIBOUTI
- DM : DOMINICA
- DO : DOMINICAN REPUBLIC
- EC : ECUADOR
- EG : EGYPT
- SV : EL SALVADOR
- GQ : EQUATORIAL GUINEA
- ER : ERITREA
- EE : ESTONIA
- ET : ETHIOPIA
- FK : FALKLAND ISLANDS (MALVINAS)
- FO : FAROE ISLANDS
- FJ : FIJI
- FI : FINLAND
- FR : FRANCE
- GF : FRENCH GUIANA
- PF : FRENCH POLYNESIA
- TF : FRENCH SOUTHERN TERRITORIES
- GA : GABON
- GM : GAMBIA
- GE : GEORGIA
- DE : GERMANY
- GH : GHANA
- GI : GIBRALTAR
- GR : GREECE
- GL : GREENLAND
- GD : GRENADA
- GP : GUADELOUPE
- GU : GUAM
- GT : GUATEMALA
- GG : GUERNSEY
- GN : GUINEA
- GW : GUINEA-BISSAU
- GY : GUYANA
- HT : HAITI
- HM : HEARD ISLAND AND MCDONALD ISLANDS
- VA : HOLY SEE (VATICAN CITY STATE)
- HN : HONDURAS
- HK : HONG KONG
- HU : HUNGARY
- IS : ICELAND
- IN : INDIA
- ID : INDONESIA
- IR : IRAN, ISLAMIC REPUBLIC OF
- IQ : IRAQ
- IE : IRELAND
- IM : ISLE OF MAN
- IL : ISRAEL
- IT : ITALY
- JM : JAMAICA
- JP : JAPAN
- JE : JERSEY
- JO : JORDAN
- KZ : KAZAKHSTAN
- KE : KENYA
- KI : KIRIBATI
- KP : KOREA, DEMOCRATIC PEOPLES REPUBLIC OF
- KR : KOREA, REPUBLIC OF
- KW : KUWAIT
- KG : KYRGYZSTAN
- LA : LAO PEOPLES DEMOCRATIC REPUBLIC
- LV : LATVIA
- LB : LEBANON
- LS : LESOTHO
- LR : LIBERIA
- LY : LIBYAN ARAB JAMAHIRIYA
- LI : LIECHTENSTEIN
- LT : LITHUANIA
- LU : LUXEMBOURG
- MO : MACAO
- MK : NORTH MACEDONIA
- MG : MADAGASCAR
- MW : MALAWI
- MY : MALAYSIA
- MV : MALDIVES
- ML : MALI
- MT : MALTA
- MH : MARSHALL ISLANDS
- MQ : MARTINIQUE
- MR : MAURITANIA
- MU : MAURITIUS
- YT : MAYOTTE
- MX : MEXICO
- FM : MICRONESIA, FEDERATED STATES OF
- MD : MOLDOVA, REPUBLIC OF
- MC : MONACO
- MN : MONGOLIA
- ME : MONTENEGRO
- MS : MONTSERRAT
- MA : MOROCCO
- MZ : MOZAMBIQUE
- MM : MYANMAR
- NA : NAMIBIA
- NR : NAURU
- NP : NEPAL
- NL : NETHERLANDS
- NC : NEW CALEDONIA
- NZ : NEW ZEALAND
- NI : NICARAGUA
- NE : NIGER
- NG : NIGERIA
- NU : NIUE
- NF : NORFOLK ISLAND
- MP : NORTHERN MARIANA ISLANDS
- NO : NORWAY
- OM : OMAN
- PK : PAKISTAN
- PW : PALAU
- PS : PALESTINIAN TERRITORY, OCCUPIED
- PA : PANAMA
- PG : PAPUA NEW GUINEA
- PY : PARAGUAY
- PE : PERU
- PH : PHILIPPINES
- PN : PITCAIRN
- PL : POLAND
- PT : PORTUGAL
- PR : PUERTO RICO
- QA : QATAR
- RE : REUNION
- RO : ROMANIA
- RU : RUSSIAN FEDERATION
- RW : RWANDA
- BL : SAINT BARTHELEMY
- SH : SAINT HELENA
- KN : SAINT KITTS AND NEVIS
- LC : SAINT LUCIA
- MF : SAINT MARTIN
- PM : SAINT PIERRE AND MIQUELON
- VC : SAINT VINCENT AND THE GRENADINES
- WS : SAMOA
- SM : SAN MARINO
- ST : SAO TOME AND PRINCIPE
- SA : SAUDI ARABIA
- SN : SENEGAL
- RS : SERBIA
- SC : SEYCHELLES
- SL : SIERRA LEONE
- SG : SINGAPORE
- SK : SLOVAKIA
- SI : SLOVENIA
- SB : SOLOMON ISLANDS
- SO : SOMALIA
- ZA : SOUTH AFRICA
- GS : SOUTH GEORGIA AND THE SOUTH SANDWICH ISLANDS
- SS : SOUTH SUDAN
- ES : SPAIN
- LK : SRI LANKA
- SD : SUDAN
- SR : SURINAME
- SJ : SVALBARD AND JAN MAYEN
- SZ : ESWATINI
- SE : SWEDEN
- CH : SWITZERLAND
- SY : SYRIAN ARAB REPUBLIC
- TW : TAIWAN, PROVINCE OF CHINA
- TJ : TAJIKISTAN
- TZ : TANZANIA, UNITED REPUBLIC OF
- TH : THAILAND
- TL : TIMOR-LESTE
- TG : TOGO
- TK : TOKELAU
- TO : TONGA
- TT : TRINIDAD AND TOBAGO
- TN : TUNISIA
- TR : TURKEY
- TM : TURKMENISTAN
- TC : TURKS AND CAICOS ISLANDS
- TV : TUVALU
- UG : UGANDA
- UA : UKRAINE
- AE : UNITED ARAB EMIRATES
- GB : UNITED KINGDOM
- US : UNITED STATES
- UM : UNITED STATES MINOR OUTLYING ISLANDS
- UY : URUGUAY
- UZ : UZBEKISTAN
- VU : VANUATU
- VE : VENEZUELA, BOLIVARIAN REPUBLIC OF
- VN : VIET NAM
- VG : VIRGIN ISLANDS, BRITISH
- VI : VIRGIN ISLANDS, U.S.
- WO : Indeterminately reserved
- WF : WALLIS AND FUTUNA
- EH : WESTERN SAHARA
- YE : YEMEN
- ZM : ZAMBIA
- ZW : ZIMBABWE


#### 3.28. customercurrency {#type.customercurrency}

3-letter Currency Codes

- EUR
- CHF
- USD


#### 3.29. date {#type.date}

Date string in the format yyyy-MM-dd


#### 3.30. dateTime {#type.datetime}

Date in ISO 8601 format


#### 3.31. dnskey {#type.dnskey}

DNSKEY presentation format.


#### 3.32. dnssecAlgorithm {#type.dnssecalgorithm}

Supported algorithms for DNSSEC.

- 5 : RSA/SHA-1
- 7 : RSASHA1-NSEC3-SHA1
- 8 : RSA/SHA-256
- 10 : RSA/SHA-512
- 13 : ECDSA Curve P-256 with SHA-256
- 14 : ECDSA Curve P-384 with SHA-384
- 15 : ED25519
- 16 : ED448


#### 3.33. dnssecDigestType {#type.dnssecdigesttype}

Supported digest types for DNSSEC.

- 2 : SHA256
- 4 : SHA384


#### 3.34. dnssecDomainStatus {#type.dnssecdomainstatus}

DNSSEC domain status values.

- DELETE_ALL : All DNSSEC keys for this domain shall be deleted.
- MANUAL : DNSSEC keys are manually managed.
- UPDATE : A manual DNSSEC update is pending.


#### 3.35. dnssecFlag {#type.dnssecflag}

Supported flag values for DNSSEC.

- 256 : ZONE
- 257 : ZONE+SEP


#### 3.36. dnssecKeyStatus {#type.dnsseckeystatus}

DNSKEY status values.

- CREATE
- DELAYED
- OK
- DELETED
- IGNORE


#### 3.37. documentformat {#type.documentformat}

Format of a requested document

- RAW
- raw
- PDF
- pdf
- CSV
- csv


#### 3.38. domainLogOrder {#type.domainlogorder}

Domain log sort order values

- LOGTIMEASC
- LOGTIMEDESC


#### 3.39. domainOrder {#type.domainorder}

Domain list sort order values

- DOMAINASC
- DOMAINDESC
- STATUSASC
- STATUSDESC
- CRDATEASC
- CRDATEDESC
- EXDATEASC
- EXDATEDESC
- REDATEASC
- REDATEDESC
- TRANSFERLOCKASC
- TRANSFERLOCKDESC
- UPDATEASC
- UPDATEDESC
- SCDATEASC
- SCDATEDESC
- RENEWALMODEASC
- RENEWALMODEDESC


#### 3.40. ds {#type.ds}

DS presentation format.


#### 3.41. dunsNumber {#type.dunsnumber}

DUNS number


#### 3.42. eea_countries {#type.eea_countries}

EEA countries in 2 letter country codes (ISO-3166-1)

- AT : AUSTRIA
- BE : BELGIUM
- BG : BULGARIA
- HR : CROATIA
- CY : CYPRUS
- CZ : CZECH REPUBLIC
- DK : DENMARK
- EE : ESTONIA
- FI : FINLAND
- FR : FRANCE
- GF : FRENCH GUIANA
- PF : FRENCH POLYNESIA
- TF : FRENCH SOUTHERN TERRITORIES
- DE : GERMANY
- GR : GREECE
- HU : HUNGARY
- IS : ICELAND
- IE : IRELAND
- IT : ITALY
- LV : LATVIA
- LI : LIECHTENSTEIN
- LT : LITHUANIA
- LU : LUXEMBOURG
- MT : MALTA
- NL : NETHERLANDS
- NO : NORWAY
- PL : POLAND
- PT : PORTUGAL
- RO : ROMANIA
- SK : SLOVAKIA
- SI : SLOVENIA
- ES : SPAIN
- SE : SWEDEN


#### 3.43. email {#type.email}

Valid email address


#### 3.44. emailoptional {#type.emailoptional}

An optional (can be an empty string) valid email address


#### 3.45. esIdType {#type.esidtype}

.ES owner ID type

- 1 : Spanish personal/tax identity number (NIF)
- 3 : Spanish foreign resident number (NIE)
- 0 : Other


#### 3.46. esLegalForm {#type.eslegalform}

.ES owner legal form

- 1 : Individual
- 877 : Others
- 39 : Economic Interest Group
- 47 : Association
- 59 : Sports Association
- 68 : Professional Association
- 124 : Savings Bank
- 150 : Community Property
- 152 : Community of Owners
- 164 : Order or Religious Institution
- 181 : Consulate
- 197 : Public Law Association
- 203 : Embassy
- 229 : Local Authority
- 269 : Sports Federation
- 286 : Foundation
- 365 : Mutual Insurance Company
- 434 : Regional Government Body
- 436 : Central Government Body
- 439 : Political Party
- 476 : Trade Union
- 510 : Farm Partnership
- 524 : Public Limited Company
- 525 : Sports Association
- 554 : Civil Society
- 560 : General Partnership
- 562 : General and Limited Partnership
- 566 : Cooperative
- 608 : Worker-owned Company
- 612 : Limited Company
- 713 : Spanish Office
- 717 : Temporary Alliance of Enterprises
- 744 : Worker-owned Limited Company
- 745 : Regional Public Entity
- 746 : National Public Entity
- 747 : Local Public Entity
- 878 : Designation of Origin Supervisory Council
- 879 : Entity Managing Natural Areas


#### 3.47. es_nif_nie {#type.es_nif_nie}

Spanish NIF/NIE


#### 3.48. extData {#type.extdata}

Domain extra data

- NO-DOMAIN-CHECK (boolean)
- NO-HOST-CHECK (boolean)
- BIRTH-DATE (date)
- BIRTH-CITY (text64)
- BIRTH-COUNTRY (country)
- BIRTH-PC (text64)
- VAT-NUMBER (vatNoInternational)
- IDCARD-OR-PASSPORT-NUMBER (text64)
- SELLER-IDCARD-OR-PASSPORT-NUMBER (text64)
- IDCARD-OR-PASSPORT-ISSUER (text64)
- IDCARD-OR-PASSPORT-ISSUE-DATE (date)
- IDCARD-OR-PASSPORT-CC (country)
- COMPANY-NUMBER (text64)
- SELLER-CEO-NAME (token255)
- SELLER-VAT-NUMBER (vatNoInternational)
- FORMATION-DATE (date)
- ACCEPT-TRUSTEE-TAC (boolean)
- WHOIS-PROTECTION (boolean)
- NICSE-IDNUMBER (seIdNo)
- IT-CODICE-FISCALE (itCodiceFiscale)
- EXPIRE-DOMAIN (boolean)
- SK-LEGAL-FORM (skLegalForm)
- CA-LEGALTYPE (caLegalType)
- NAME-EMAIL-FORWARD (email)
- AE-REGISTRANT-WARRANTY-STATEMENT (_true)
- TRAVEL-IN-INDUSTRY (_true)
- US-NEXUS-APPPURPOSE (usPurpose)
- US-NEXUS-CATEGORY (usCategory)
- CN-REGISTRATION-TAC-APPROVED (boolean)
- CN-ICP-NUMBER (text64)
- TAX-NUMBER (text64)
- TAX-OFFICE (text64)
- TECH-VAT-NUMBER (vatNoInternational)
- TECH-TAX-NUMBER (text64)
- TRADEMARK-NAME (text64)
- TRADEMARK-NUMBER (text64)
- TRADEMARK-DATE (date)
- FI-HENKILOTUNNUS (fiHenkilotunnus)
- IE-HOLDER-TYPE (ieHolderType)
- PT-LEGITIMACY (ptLegitimacy)
- PT-REGISTRATION-BASIS (ptRegistrationBasis)
- AERO-ENS-AUTH-ID (text64)
- AERO-ENS-AUTH-KEY (text64)
- LEGAL-REPRESENTATIVE-POSITION (text64)
- ADMIN-BIRTH-DATE (date)
- ADMIN-IDCARD-OR-PASSPORT-NUMBER (text64)
- TECH-BIRTH-DATE (date)
- TECH-IDCARD-OR-PASSPORT-NUMBER (text64)
- AU-DOMAIN-IDTYPE (auDomainIdType)
- AU-DOMAIN-RELATION (auDomainRelation)
- AU-DOMAIN-RELATIONTYPE (auDomainRelationType)
- AU-ELIGIBILITY-ID-TYPE (auEligibilityIdType)
- DUNS-NUMBER (dunsNumber)
- EEA-LOCAL-ID (text64)
- UA-TRADEMARK-TYPE (uaTrademarkType)
- KR-CTFY-TYPE (krCtfyType)
- KR-CTFY-NO (text64)
- HR-OIB (hrOib)
- OWNER-CONTACT-NAME-IS-LEGAL-REPRESENTATIVE (_true)
- BG-REGISTRATION-TAC-APPROVED (_true)
- BG-APPLICATION-NUMBER (bgAppNumber)
- BG-APPLICATION-DATE (date)
- IR-NATIONAL-ID (irNationalId)
- IR-ORGANIZATION-ID (irOrganizationId)
- IR-COMPANY-REGISTRATION-CC (country)
- IR-COMPANY-REGISTRATION-SP (text64)
- IR-COMPANY-REGISTRATION-CENTER (text64)
- IR-COMPANY-REGISTRATION-TYPE (irCompanyRegistrationType)
- BR-CPF (brCpf)
- BR-CNPJ (brCnpj)
- RU-KPP (text64)
- SG-ADMIN-RCBID (text64)
- SG-ADMIN-SINGPASSID (text64)
- TR-CITIZEN-ID (trCitizenId)
- COMTR-REGISTRATION-TAC-APPROVED (_true)
- NO-PERSON-IDENTIFIER (noPersonIdentifier)
- IS-KENNITALA (isKennitala)
- HK-INDUSTRY-TYPE (hkIndustryType)
- MY-ORG-TYPE (myOrgType)
- ES-LEGAL-FORM (esLegalForm)
- ES-ID-TYPE (esIdType)
- ES-NIF-NIE (es_nif_nie)
- ES-ADMIN-ID-TYPE (esIdType)
- ES-ADMIN-NIF-NIE (es_nif_nie)
- ES-ID (text64)
- ES-ADMIN-ID (text64)
- ES-TECH-ID (text64)
- ES-BILLING-ID (text64)
- INTENDED-USE (text255)
- ALLOCATION-TOKEN (text255)
- LU-TRADE-TAC-APPROVED (text255)
- LU-TRANSFER-TAC-APPROVED (text255)
- SELLER-EMAIL (email)
- ACKNOWLEDGE-SECURE-ONLY-APP (_true)
- ACKNOWLEDGE-SECURE-ONLY-PAGE (_true)
- ACKNOWLEDGE-SECURE-ONLY-DEV (_true)
- ACKNOWLEDGE-SECURE-ONLY-NEW (_true)
- ACKNOWLEDGE-SECURE-ONLY-DAY (_true)
- ACKNOWLEDGE-SECURE-ONLY-BOO (_true)
- ACKNOWLEDGE-SECURE-ONLY-RSVP (_true)
- ACKNOWLEDGE-SECURE-ONLY-DAD (_true)
- ACKNOWLEDGE-SECURE-ONLY-ESQ (_true)
- ACKNOWLEDGE-SECURE-ONLY-FOO (_true)
- ACKNOWLEDGE-SECURE-ONLY-NEXUS (_true)
- ACKNOWLEDGE-SECURE-ONLY-PROF (_true)
- ACKNOWLEDGE-SECURE-ONLY-ZIP (_true)
- ACKNOWLEDGE-SECURE-ONLY-MOV (_true)
- ACKNOWLEDGE-SECURE-ONLY-PHD (_true)
- ACKNOWLEDGE-SECURE-ONLY-MEME (_true)
- ACKNOWLEDGE-SECURE-ONLY-ING (_true)
- ACKNOWLEDGE-SECURE-ONLY-CHANNEL (_true)
- CHANNEL-ACCEPT-REQUIREMENTS (_true)
- DOT-NEW-CONFIRMATION (_true)
- LEGAL-REPRESENTATIVE (text255)
- EU-COUNTRY-OF-CITIZENSHIP (eea_countries)
- GAY-ACCEPT-REQUIREMENTS (_true)
- COOP-ACCEPT-REQUIREMENTS (_true)
- NGO-ACCEPT-REQUIREMENTS (_true)
- MUSIC-ACCEPT-REQUIREMENTS (_true)
- BANK-ACCEPT-REQUIREMENTS (_true)
- INSURANCE-ACCEPT-REQUIREMENTS (_true)
- GM-ACCEPT-REQUIREMENTS (_true)
- CA-ACCEPT-REGISTRANT-AGREEMENT (_true)
- SWISS-UPI (swissupi)
- SWISS-UID (swissuid)
- ZUERICH-UID (zuerichuid)
- DK-ACCEPT-REGISTRAR (_true)
- DK-ACCEPT-TAC (_true)


#### 3.49. fiHenkilotunnus {#type.fihenkilotunnus}

FI Henkilötunnus


#### 3.50. float {#type.float}

A valid float value


#### 3.51. float_signed {#type.float_signed}

A valid signed float value


#### 3.52. hkIndustryType {#type.hkindustrytype}

In which industry does a company owning a .hk domain operate in?

- 0 : None
- 010100 : Plastics, Petro-Chemicals, Chemicals - Plastics & Plastic Products
- 010200 : Plastics, Petro-Chemicals, Chemicals - Rubber & Rubber Products
- 010300 : Plastics, Petro-Chemicals, Chemicals - Fibre Materials & Products
- 010400 : Plastics, Petro-Chemicals, Chemicals - Petroleum, Coal & Other Fuels
- 010500 : Plastics, Petro-Chemicals, Chemicals - Chemicals & Chemical Products
- 020100 : Metals, Machinery, Equipment - Metal Materials & Treatment
- 020200 : Metals, Machinery, Equipment - Metal Products
- 020300 : Metals, Machinery, Equipment - Industrial Machinery & Supplies
- 020400 : Metals, Machinery, Equipment - Precision & Optical Equipment
- 020500 : Metals, Machinery, Equipment - Moulds & Dies
- 030100 : Printing, Paper, Publishing - Printing, Photocopying, Publishing
- 030200 : Printing, Paper, Publishing - Paper, Paper Products
- 040100 : Construction, Decoration, Environmental Engineering - Construction Contractors
- 040200 : Construction, Decoration, Environmental Engineering - Construction Materials
- 040300 : Construction, Decoration, Environmental Engineering - Decoration Materials
- 040400 : Construction, Decoration, Environmental Engineering - Construction, Safety Equipment & Supplies
- 040500 : Construction, Decoration, Environmental Engineering - Decoration, Locksmiths, Plumbing & Electrical Works
- 040600 : Construction, Decoration, Environmental Engineering - Fire Protection Equipment & Services
- 040700 : Construction, Decoration, Environmental Engineering - Environmental Engineering, Waste Reduction
- 050100 : Textiles, Clothing & Accessories - Textiles, Fabric
- 050200 : Textiles, Clothing & Accessories - Clothing
- 050300 : Textiles, Clothing & Accessories - Uniforms, Special Clothing
- 050400 : Textiles, Clothing & Accessories - Clothing Manufacturing Accessories
- 050500 : Textiles, Clothing & Accessories - Clothing Processing & Equipment
- 050600 : Textiles, Clothing & Accessories - Fur, Leather & Leather Goods
- 050700 : Textiles, Clothing & Accessories - Handbags, Footwear, Optical Goods, Personal Accessories
- 060100 : Electronics, Electrical Appliances - Electronic Equipment & Supplies
- 060200 : Electronics, Electrical Appliances - Electronic Parts & Components
- 060300 : Electronics, Electrical Appliances - Electrical Appliances, Audio-Visual Equipment
- 070100 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Kitchenware, Tableware
- 070200 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Bedding
- 070300 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Bathroom, Cleaning Accessories
- 070400 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Household Goods
- 070500 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Wooden, Bamboo & Rattan Goods
- 070600 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Home Furnishings, Arts & Crafts
- 070700 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Watches, Clocks
- 070800 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Jewellery Accessories
- 070900 : Houseware, Watches, Clocks, Jewellery, Toys, Gifts - Toys, Games, Gifts
- 080100 : Business & Professional Services, Finance - Accounting, Legal Services
- 080200 : Business & Professional Services, Finance - Advertising, Promotion Services
- 080300 : Business & Professional Services, Finance - Consultancy Services
- 080400 : Business & Professional Services, Finance - Translation, Design Services
- 080500 : Business & Professional Services, Finance - Cleaning, Pest Control Services
- 080600 : Business & Professional Services, Finance - Security Services
- 080700 : Business & Professional Services, Finance - Trading, Business Services
- 080800 : Business & Professional Services, Finance - Employment Services
- 080900 : Business & Professional Services, Finance - Banking, Finance, Investment
- 081000 : Business & Professional Services, Finance - Insurance
- 081100 : Business & Professional Services, Finance - Property, Real Estate
- 090100 : Transportation, Logistics - Land Transport, Motorcars
- 090200 : Transportation, Logistics - Sea Transport, Boats
- 090300 : Transportation, Logistics - Air Transport
- 090400 : Transportation, Logistics - Moving, Warehousing, Courier & Logistics Services
- 090500 : Transportation, Logistics - Freight Forwarding
- 100100 : Office Equipment, Furniture, Stationery, Information Technology - Office, Commercial Equipment & Supplies
- 100200 : Office Equipment, Furniture, Stationery, Information Technology - Office & Home Furniture
- 100300 : Office Equipment, Furniture, Stationery, Information Technology - Stationery & Educational Supplies
- 100400 : Office Equipment, Furniture, Stationery, Information Technology - Telecommunication Equipment & Services
- 100500 : Office Equipment, Furniture, Stationery, Information Technology - Computers, Information Technology
- 110100 : Food, Flowers, Fishing & Agriculture - Food Products & Supplies
- 110200 : Food, Flowers, Fishing & Agriculture - Beverages, Tobacco
- 110300 : Originál Food, Flowers, Fishing & Agriculture - Restaurant Equipment & Supplies
- 110400 : Food, Flowers, Fishing & Agriculture - Flowers, Artificial Flowers, Plants
- 110500 : Food, Flowers, Fishing & Agriculture - Fishing
- 110600 : Food, Flowers, Fishing & Agriculture - Agriculture
- 120100 : Medical Services, Beauty, Social Services - Medicine & Herbal Products
- 120200 : Medical Services, Beauty, Social Services - Medical & Therapeutic Services
- 120300 : Medical Services, Beauty, Social Services - Medical Equipment & Supplies
- 120400 : Medical Services, Beauty, Social Services - Beauty, Health
- 120500 : Medical Services, Beauty, Social Services - Personal Services
- 120600 : Medical Services, Beauty, Social Services - Organizations, Associations
- 120700 : Medical Services, Beauty, Social Services - Information, Media
- 120800 : Medical Services, Beauty, Social Services - Public Utilities
- 120900 : Medical Services, Beauty, Social Services - Religion, Astrology, Funeral Services
- 130100 : Culture, Education - Music, Arts
- 130200 : Culture, Education - Learning Instruction & Training
- 130300 : Culture, Education - Elementary Education
- 130400 : Culture, Education - Tertiary Education, Other Education Services
- 130500 : Culture, Education - Sporting Goods
- 130600 : Culture, Education - Sporting, Recreational Facilities & Venues
- 130700 : Culture, Education - Hobbies, Recreational Activities
- 130800 : Culture, Education - Pets, Pets Services & Supplies
- 140101 : Dining, Entertainment, Shopping, Travel - Restaurant Guide - Chinese
- 140102 : Dining, Entertainment, Shopping, Travel - Restaurant Guide - Asian
- 140103 : Dining, Entertainment, Shopping, Travel - Restaurant Guide - Western
- 140200 : Dining, Entertainment, Shopping, Travel - Catering Services, Eateries
- 140300 : Dining, Entertainment, Shopping, Travel - Entertainment Venues
- 140400 : Dining, Entertainment, Shopping, Travel - Entertainment Production & Services
- 140500 : Dining, Entertainment, Shopping, Travel - Entertainment Equipment & Facilities
- 140600 : Dining, Entertainment, Shopping, Travel - Shopping Venues
- 140700 : Dining, Entertainment, Shopping, Travel - Travel, Hotels & Accommodation


#### 3.53. hostname {#type.hostname}

A valid hostname


#### 3.54. hrOib {#type.hroib}

Croatian personal identification number


#### 3.55. ieHolderType {#type.ieholdertype}

.IE Holder Type

- COMPANY : Company (Ltd., PLC, etc.)
- CHARITY : Charity
- OTHER : Natural Person/Other


#### 3.56. int {#type.int}

A valid integer value


#### 3.57. ip {#type.ip}

A valid ip (v4 or v6) address


#### 3.58. ipList {#type.iplist}

A valid list of ip (v4 or v6) addresses


#### 3.59. ip_url {#type.ip_url}

A valid ip (v4 or v6) or url address


#### 3.60. irCompanyRegistrationType {#type.ircompanyregistrationtype}

Company registration type

- PublicCompany : Public Company
- PrivateCompany : Private Company
- LimitedCompany : Limited Company
- CooperativeCompany : Cooperative Company
- Organization : Institute
- PressAndPublication : Press & Publication


#### 3.61. irNationalId {#type.irnationalid}

Iranian National ID


#### 3.62. irOrganizationId {#type.irorganizationid}

Iranian Organization ID


#### 3.63. isKennitala {#type.iskennitala}

Icelandic ID number (Kennitala)


#### 3.64. itCodiceFiscale {#type.itcodicefiscale}

Italian Fiscal Code


#### 3.65. krCtfyType {#type.krctfytype}

.KR Certify Type

- BUSINESS : Business Registration Certificate
- SOCIAL : Social Security Card
- CORP : Corporation Registration Certificate
- UNIQUE : Unique Number Certificate
- TAX : Tax Registration Certificate
- ORG : Organization Registration Certificate
- BRAND : Brand Name Registration Certificate
- SERVICE : Service Name Registration Certificate
- SCHOOL : School Foundation Certificate
- AUXLAB : Auxiliary LAB Certificate
- ORGVOU : Organization Voucher
- ESTABLISHMENT : Establishment Authorization Certificate
- BYLAWS : Bylaws & Rules
- FOREIGNER : Foreigner registration Certificate
- DRIVELIC : Drive License
- PASSPORT : Passport
- ETC : etc


#### 3.66. language {#type.language}

Language in 2 letter code

- DE : Deutsch
- EN : English
- ES : Español


#### 3.67. message_status {#type.message_status}

Possible list of domain status

- APP FAILED
- APP INITIATED
- APP REQUESTED
- APP SUCCESSFUL
- AUTO-CLOSE NOTIFICATION
- CONTACTVERIFICATION NOTIFY
- CONTACTVERIFICATION REACTIVATED
- CONTACTVERIFICATION SUSPENDED
- DELETE CANCELED
- DELETE DELAYED
- DELETE FAILED
- DELETE INITIATED
- DELETE NOTIFY
- DELETE REQUESTED
- DELETE SCHEDULED
- DELETE SUCCESSFUL
- DNSSEC DEACTIVATION FAILED
- DNSSEC DEACTIVATION REQUESTED
- DNSSEC DEACTIVATION SUCCESSFUL
- EXPIRED
- EXPIRE CANCELED
- EXPIRE DELAYED
- EXPIRE INITIATED
- EXPIRE NOTIFICATION
- EXPIRE REQUESTED
- EXPIRE SCHEDULED
- EXPIRE SUCCESSFUL
- PUSH CANCELED
- PUSH DELAYED
- PUSH FAILED
- PUSH INITIATED
- PUSH REQUESTED
- PUSH SCHEDULED
- PUSH SUCCESSFUL
- REG CANCELED
- REG CHANGED
- REG DELAYED
- REG FAILED
- REG INITIATED
- REG NOTIFY
- REG REQUESTED
- REG SCHEDULED
- REG SUCCESSFUL
- REG WAIT
- REG WAITING
- RENEWAL CANCELED
- RENEWAL DELAYED
- RENEWAL FAILED
- RENEWAL INITIATED
- RENEWAL REQUESTED
- RENEWAL SUCCESSFUL
- RENEWALMODE AUTORENEW
- SAFETY CANCELED
- SAFETY DELAYED
- SAFETY INITIATED
- SAFETY REQUESTED
- SAFETY SUCCESSFUL
- STATUSCHANGE DELAYED
- STATUSCHANGE SUCCESSFUL
- TRADE DELAYED
- TRADE FAILED
- TRADE INITIATED
- TRADE NOTIFY
- TRADE REQUESTED
- TRADE SCHEDULED
- TRADE STARTED
- TRADE SUCCESSFUL
- TRADE WAITING
- TRADE-OUT SUCCESSFUL
- TRANSFER CANCELED
- TRANSFER CHANGED
- TRANSFER DELAYED
- TRANSFER FAILED
- TRANSFER INITIATED
- TRANSFER NOTIFY
- TRANSFER PENDING
- TRANSFER REQUESTED
- TRANSFER SCHEDULED
- TRANSFER STARTED
- TRANSFER SUCCESSFUL
- TRANSFER WAIT
- TRANSFER WAITING
- TRANSFER-OUT ACK
- TRANSFER-OUT FAILED
- TRANSFER-OUT NACK
- TRANSFER-OUT NOTIFY
- TRANSFER-OUT REQUESTED
- TRANSFER-OUT SCHEDULED
- TRANSFER-OUT SUCCESSFUL
- TRANSFERLOCK ACTIVATION
- TRANSFERLOCK DEACTIVATION
- TRANSFERLOCK SCHEDULED
- UNDELETE CANCELED
- UNDELETE DELAYED
- UNDELETE INITIATED
- UNDELETE REQUESTED
- UNDELETE SUCCESSFUL
- UNEXPIRE REQUESTED
- UNEXPIRE SUCCESSFUL
- UPDATE CANCELED
- UPDATE CHANGED
- UPDATE DELAYED
- UPDATE FAILED
- UPDATE INITIATED
- UPDATE NOTIFY
- UPDATE REQUESTED
- UPDATE SCHEDULED
- UPDATE STARTED
- UPDATE SUCCESSFUL
- USERTRANSFER APPROVED
- USERTRANSFER CANCELED
- USERTRANSFER DENIED
- USERTRANSFER FAILED
- USERTRANSFER STARTED
- USERTRANSFER SUCCESSFUL
- USERTRANSFER-OUT APPROVED
- USERTRANSFER-OUT CANCELED
- USERTRANSFER-OUT DENIED
- USERTRANSFER-OUT FAILED
- USERTRANSFER-OUT REQUESTED
- USERTRANSFER-OUT SUCCESSFUL
- WHOIS REMINDER


#### 3.68. message_type {#type.message_type}

An ENUM value that describes message type

- DOMAIN
- CERTIFICATE
- CONTACT
- HOST


#### 3.69. myOrgType {#type.myorgtype}

What type of company is registering a .my domain?

- company pursuant to companies act(roc) : company pursuant to companies act(roc)
- business pursuant to business registration act(rob) : business pursuant to business registration act(rob)
- business pursuant to commercial license ordinance : business pursuant to commercial license ordinance
- architect firm : architect firm
- audit firm : audit firm
- law firm : law firm
- valuer, appraiser and estate agent firm : valuer, appraiser and estate agent firm
- offshore company : offshore company
- educational institution accredited/registered by relevant government department/agency : educational institution accredited/registered by relevant government department/agency
- farmers organisation : farmers organisation
- federal government department or agency : federal government department or agency
- foreign embassy : foreign embassy
- foreign office : foreign office
- government aided primary and/or secondary school : government aided primary and/or secondary school
- lembaga (board) : lembaga (board)
- local authority department or agency : local authority department or agency
- maktab rendah sains mara (mrsm) under the administration of mara : Nmaktab rendah sains mara (mrsm) under the administration of maraone
- ministry of defences department or agency : ministry of defences department or agency
- parents teachers association : parents teachers association
- polytechnic under ministry of education administration : polytechnic under ministry of education administration
- private higher educational institution : private higher educational institution
- private school : private school
- regional office : regional office
- religious entity : religious entity
- representative office : representative office
- society pursuant to societies act(ros) : society pursuant to societies act(ros)
- sports organisation : sports organisation
- state government department or agency : state government department or agency
- trade union : trade union
- trustee : trustee
- university under the administration of ministry of education : university under the administration of ministry of education


#### 3.70. noPersonIdentifier {#type.nopersonidentifier}

.NO Person Identifier


#### 3.71. nsList {#type.nslist}

List of valid nameserver


#### 3.72. nsSetType {#type.nssettype}

Type of nameserver set

- PRIMARY
- SECONDARY
- EXTERNAL


#### 3.73. nsType {#type.nstype}

Type of nameserver

- MASTER
- SLAVE


#### 3.74. password {#type.password}

A string value with a maximum length of 99999 characters


#### 3.75. paymentType {#type.paymenttype}

Type of payment

- PAYPAL
- BANKTRANSFER


#### 3.76. period {#type.period}

Allowed periods

- 1M
- 1Y
- 2Y
- 3Y
- 4Y
- 5Y
- 6Y
- 7Y
- 8Y
- 9Y
- 10Y


#### 3.77. phone {#type.phone}

A valid phone number in international format +49.178645376-78.


#### 3.78. phoneoptional {#type.phoneoptional}

A valid phone number in international format +49.178645376-78 (can be empty).


#### 3.79. ptLegitimacy {#type.ptlegitimacy}

.PT Legitimacy

- PC : PC - Corporate Entity
- PS : PS - Individual Entity


#### 3.80. ptRegistrationBasis {#type.ptregistrationbasis}

.PT Registration Basis

- L : L - Domain will pre-registered for liberalization of .PT on May 1, 2012
- 7A : 7A - Coincides with constituted right owned by the applicant
- 6A : 6A - Coincides with the name, abbreviation or acronym of the domain holder
- 6J : 6J - Coincides with the name, abbreviation or acronym of a Self-employed Person
- 6K : 6K - Coincides with the document as Liberal Professional


#### 3.81. recordType {#type.recordtype}

Type of record

- A
- AAAA
- AFSDB
- ALIAS
- CAA
- CERT
- CNAME
- HINFO
- HTTPS
- IPSECKEY
- LOC
- MX
- NAPTR
- NS
- OPENPGPKEY
- PTR
- RP
- SMIMEA
- SOA
- SRV
- SSHFP
- SVCB
- TLSA
- TXT
- URI
- URL


#### 3.82. remarks {#type.remarks}

A string value with a maximum length of 255 characters


#### 3.83. renewalMode {#type.renewalmode}

Allowed renewal modes

- AUTORENEW
- AUTODELETE
- AUTOEXPIRE


#### 3.84. seIdNo {#type.seidno}

SE Id-Number


#### 3.85. signMethod {#type.signmethod}

How the signature is delivered

- WEB


#### 3.86. skLegalForm {#type.sklegalform}

.SK Legal Form

- AS : Company (a.s.)
- FO : Personal
- OTHER : Other
- SRO : Company (s.r.o.)
- Z : self employed


#### 3.87. swissuid {#type.swissuid}

A valid Swiss UID number


#### 3.88. swissupi {#type.swissupi}

A valid Swiss UPI (Universal Person Identification) number


#### 3.89. tagUpdateAdd {#type.tagupdateadd}

atomic tag update

- domainId (array_int)


#### 3.90. tagUpdateRem {#type.tagupdaterem}

atomic tag update

- domainId (array_int)


#### 3.91. text {#type.text}

A string value with a minimum length of 1 and a maximum length of 10000 characters


#### 3.92. text0 {#type.text0}

A string value with a maximum length of 10000 characters


#### 3.93. text0100 {#type.text0100}

A string value with a maximum length of 100 characters


#### 3.94. text0255 {#type.text0255}

A string value with newline characters and a maximum length of 255 characters


#### 3.95. text064 {#type.text064}

A string value with a maximum length of 64 characters


#### 3.96. text10 {#type.text10}

A string value with a minimum length of 2 and a maximum length of 10 characters


#### 3.97. text1024 {#type.text1024}

A string value with a minimum length of 1 and a maximum length of 1024 characters


#### 3.98. text255 {#type.text255}

A string value with newline characters, a minimum length of 1 and a maximum length of 255 characters


#### 3.99. text64 {#type.text64}

A string value with a minimum length of 1 and a maximum length of 64 characters


#### 3.100. tfaMethod {#type.tfamethod}

The two factor method that is used or the string 0 if none is used

- 0
- GOOGLE-AUTH


#### 3.101. timestamp {#type.timestamp}

Timestamp in the ISO-8601 format (e.g. '2004-05-23T14:25:10')


#### 3.102. token0255 {#type.token0255}

A string value without newline characters and a maximum length of 255 characters


#### 3.103. token255 {#type.token255}

A string value without new line characters, with a minimum length of 1 and a maximum length of 255 characters


#### 3.104. trCitizenId {#type.trcitizenid}

Turkish Identification Number


#### 3.105. transferAnswer {#type.transferanswer}

Transfer answer

- ACK
- NACK


#### 3.106. transferMode {#type.transfermode}

Allowed transfer modes

- DEFAULT
- AUTOAPPROVE
- AUTODENY


#### 3.107. uaTrademarkType {#type.uatrademarktype}

.UA Trademark Type

- UA : Ukrainian Institute of Industrial Property (UIIP)
- WO : World Intellectual Property Organization (OMPI/MOIP/WIPO)


#### 3.108. urlRedirectType {#type.urlredirecttype}

Allowed types of url redirection

- HEADER301
- HEADER302
- FRAME


#### 3.109. usCategory {#type.uscategory}

US Nexus Category

- C11 : A natural person who is a United States citizen
- C12 : A natural person who is a permanent resident of the United States of America, or any of its possessions or territories
- C21 : A U.S.-based organization or company [A U.S.-based organization or company formed within one of the fifty (50) U.S. states, the District of Columbia, or any of the United States possessions or territories, or orga nized or otherwise constituted under the laws of a state of the United States of America, the District of Columbia or any of its possessions or territories or a U.S. federal, state, or local government entity or a political subdivision thereof.]
- C31 : A foreign entity or organization [A foreign entity or organization that has a bona fide presence in the United States of America or any of its possessions or territories who regularly engages in lawful activities (sales of goods or services or other business, commercial or non-commercial, including not-for-profit relations in the United States).]
- C32 : Entity has an office or other facility in the United States


#### 3.110. usPurpose {#type.uspurpose}

US Nexus Application Purpose

- P1 : Business use for profit
- P2 : Non-profit business, club, association, religious organization, etc.
- P3 : Personal use
- P4 : Education purposes
- P5 : Government purposes


#### 3.111. username {#type.username}

A valid username (between 5 and 42 characters)


#### 3.112. vatNo {#type.vatno}

A valid VAT number


#### 3.113. vatNoInternational {#type.vatnointernational}

A valid international VAT number


#### 3.114. zuerichuid {#type.zuerichuid}

A valid Swiss UID number or "public"


## Chapter 4: Result Codes


## Chapter 4. Result Codes {#errorcodes}

