<users>
  <user type="admin">
    <name>Elliot</name>
    <social>
      <facebook>https://facebook.com</facebook>
      <twitter>https://twitter.com</twitter>
      <youtube>https://youtube.com</youtube>
    </social>
  </user>
  <user type="reader">
    <name>Fraser</name>
    <social>
      <facebook>https://facebook.com</facebook>
      <twitter>https://twitter.com</twitter>
      <youtube>https://youtube.com</youtube>
    </social>
  </user>
</users>

///////////////////////////////////////////////////
byteValue, _ := ioutil.ReadAll(xmlFile)

var users Users

xml.Unmarshal(byteValue, &users)