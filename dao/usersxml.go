package dao

const UsersXml = `

<?xml version="1.0" encoding="UTF-8" ?>
<mapper>

	<sql id="pageImport">

	</sql>

	<select id="page1">
		select * from a where 
		<include refid="pageImport"></include>
	</select>

	<update id="xxx">
    </update>

	<delete id="xxx">
    </delete>

	<insert id="xxx">
    </insert>
</mapper>

`
