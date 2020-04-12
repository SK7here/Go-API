# MVC API using GO

## USER REFERENCE  
Run the ***Go_MVC_main.go*** file
<table border="0">
  <tr>
    <th> Task </th>
    <th> URL </th>
    <th> Method </th>
    <th> Response Code </th>
    <th> Response </th>
  </tr>
  <tr>
     <td> Check API health </td>
    <td> /SK7server/ </td>
    <td> GET </td>
    <td> 200 </td>
    <td> API alive </td>
  </tr>
    <tr>
    <td> Insert an entry </td>
    <td> /SK7server/insert/ </td>
    <td> POST </td>
    <td> 201 </td>
    <td> {Status} </td>
  </tr>
    <tr>
    <td> Read All Entries </td>
    <td> /SK7server/readAll </td>
    <td> GET </td>
    <td> 200 </td>
    <td> json of all entries </td>
  </tr>
    <tr>
    <td> Read entries by id </td>
    <td> /SK7server/readbyid/?id=value </td>
    <td> GET </td>
    <td> 200 </td>
    <td> json of entries that match id
  </tr>
    <tr>
    <td> Delete entries by id </td>
    <td> /SK7server/delete/id </td>
    <td> DELETE </td>
    <td> 200 </td>
    <td> {Status}
  </tr>
 </table>
 <br>
     

#### Sample of the table used for illustration
<table border="0">
  <tr>
    <th> name </th>
    <th> todo </th>
  </tr>
  <tr>
    <td> Kailash </td>
    <td> First operation </td>
  </tr>
  <tr>
    <td> SK7</td>
    <td> Second operation </td>
  </tr>
  
  
