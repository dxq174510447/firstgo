package dao

const UsersXml = `

<?xml version="1.0" encoding="UTF-8" ?>
<mapper>
	<select id="Find1">
			select * from users where status = 1
	</select>
	<select id="find2">
			select * from users where status = #{status} and name = #{name}
	</select>
	<select id="find3">
			select * from users where status = #{status} and name = #{name}
	</select>
	<select id="find4">
			select * from users where id = #{Id} and status = #{Status}
	</select>
	<select id="find5">
			select * from users where id = #{user.Id} and status = #{user.Status} and status = #{status}
	</select>
	<select id="find6">
			select * from users where id = #{user.Id} and status = #{user1.Status}
	</select>

</mapper>

`
