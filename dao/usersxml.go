package dao

const UsersXml = `

<?xml version="1.0" encoding="UTF-8" ?>
<mapper>
	<select id="Find1">
			select * from users 
			where status = 1
	</select>
	<select id="Find2">
			select * from users where  status = #{status}
	</select>
	<select id="Find3">
			select * from users where status = #{status} and name = #{name}
	</select>
	<select id="Find4">
			select * from users where id = #{Id} and status = #{Status}
	</select>
	<select id="Find5">
			select * from users where id > 10 order by id desc
	</select>
	<select id="Find6">
			select * from users where 
			id = #{user.Id} 
 			and status = #{user1.Status}
	</select>

</mapper>

`
